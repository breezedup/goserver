package action

import (
	"errors"
	"strconv"
	"time"

	"github.com/breezedup/goserver/core/builtin/protocol"
	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/netlib"
	"github.com/breezedup/goserver/core/transact"
	"google.golang.org/protobuf/proto"
)

type TxStartPacketFactory struct {
}

type TxStartHandler struct {
}

func (this *TxStartPacketFactory) CreatePacket() interface{} {
	pack := &protocol.TransactStart{}
	return pack
}

func (this *TxStartHandler) Process(session *netlib.Session, packetid int, data interface{}) error {
	//logger.Logger.Trace("TxStartHandler.Process")
	if ts, ok := data.(*protocol.TransactStart); ok {
		netptnp := ts.GetParenTNP()
		if netptnp == nil {
			return nil
		}
		netmtnp := ts.GetMyTNP()
		if netmtnp == nil {
			return nil
		}

		ptnp := &transact.TransNodeParam{
			TId:        transact.TransNodeID(netptnp.GetTransNodeID()),
			Tt:         transact.TransType(netptnp.GetTransType()),
			Ot:         transact.TransOwnerType(netptnp.GetOwnerType()),
			Tct:        transact.TransactCommitPolicy(netptnp.GetTransCommitType()),
			Oid:        int(netptnp.GetOwnerID()),
			SkeletonID: int(netptnp.GetSkeletonID()),
			LevelNo:    int(netptnp.GetLevelNo()),
			AreaID:     int(netptnp.GetAreaID()),
			TimeOut:    time.Duration(netptnp.GetTimeOut()),
			ExpiresTs:  netptnp.GetExpiresTs(),
		}
		mtnp := &transact.TransNodeParam{
			TId:        transact.TransNodeID(netmtnp.GetTransNodeID()),
			Tt:         transact.TransType(netmtnp.GetTransType()),
			Ot:         transact.TransOwnerType(netmtnp.GetOwnerType()),
			Tct:        transact.TransactCommitPolicy(netmtnp.GetTransCommitType()),
			Oid:        int(netmtnp.GetOwnerID()),
			SkeletonID: int(netmtnp.GetSkeletonID()),
			LevelNo:    int(netmtnp.GetLevelNo()),
			AreaID:     int(netmtnp.GetAreaID()),
			TimeOut:    time.Duration(netmtnp.GetTimeOut()),
			ExpiresTs:  netmtnp.GetExpiresTs(),
		}

		if !transact.ProcessTransStart(ptnp, mtnp, ts.GetCustomData(), mtnp.TimeOut) {
			return errors.New("TxStartHandler error, tid=" + strconv.FormatInt(netmtnp.GetTransNodeID(), 16))
		}
	}
	return nil
}

func init() {
	netlib.RegisterHandler(int(protocol.CoreBuiltinPacketID_PACKET_SS_TX_START), &TxStartHandler{})
	netlib.RegisterFactory(int(protocol.CoreBuiltinPacketID_PACKET_SS_TX_START), &TxStartPacketFactory{})
}

func ContructTxStartPacket(parent, me *transact.TransNodeParam, ud interface{}) proto.Message {
	packet := &protocol.TransactStart{
		MyTNP: &protocol.TransactParam{
			TransNodeID:     int64(me.TId),
			TransType:       int32(me.Tt),
			OwnerType:       int32(me.Ot),
			TransCommitType: int32(me.Tct),
			OwnerID:         int32(me.Oid),
			SkeletonID:      int32(me.SkeletonID),
			LevelNo:         int32(me.LevelNo),
			AreaID:          int32(me.AreaID),
			TimeOut:         int64(me.TimeOut),
		},
		ParenTNP: &protocol.TransactParam{
			TransNodeID:     int64(parent.TId),
			TransType:       int32(parent.Tt),
			OwnerType:       int32(parent.Ot),
			TransCommitType: int32(parent.Tct),
			OwnerID:         int32(parent.Oid),
			SkeletonID:      int32(parent.SkeletonID),
			LevelNo:         int32(parent.LevelNo),
			AreaID:          int32(parent.AreaID),
			TimeOut:         int64(parent.TimeOut),
		},
	}

	if ud != nil {
		b, err := netlib.MarshalPacketNoPackId(ud)
		if err != nil {
			logger.Logger.Warn("ContructTxStartPacket Marshal UserData error:", err)
		} else {
			packet.CustomData = b
		}
	}
	return packet
}
