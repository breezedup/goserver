package action

import (
	"errors"
	"strconv"

	"github.com/breezedup/goserver/core/builtin/protocol"
	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/netlib"
	"github.com/breezedup/goserver/core/transact"
	"google.golang.org/protobuf/proto"
)

type TxResultPacketFactory struct {
}

type TxResultHandler struct {
}

func (this *TxResultPacketFactory) CreatePacket() interface{} {
	pack := &protocol.TransactResult{}
	return pack
}

func (this *TxResultHandler) Process(session *netlib.Session, packetid int, data interface{}) error {
	//logger.Logger.Trace("TxResultHandler.Process")
	if tr, ok := data.(*protocol.TransactResult); ok {
		if !transact.ProcessTransResult(transact.TransNodeID(tr.GetMyTId()), transact.TransNodeID(tr.GetChildTId()), int(tr.GetRetCode()), tr.GetCustomData()) {
			return errors.New("TxResultHandler error, tid=" + strconv.FormatInt(tr.GetMyTId(), 16))
		}
	}
	return nil
}

func init() {
	netlib.RegisterHandler(int(protocol.CoreBuiltinPacketID_PACKET_SS_TX_RESULT), &TxResultHandler{})
	netlib.RegisterFactory(int(protocol.CoreBuiltinPacketID_PACKET_SS_TX_RESULT), &TxResultPacketFactory{})
}

func ContructTxResultPacket(parent, me *transact.TransNodeParam, tr *transact.TransResult) proto.Message {
	packet := &protocol.TransactResult{
		MyTId:    int64(parent.TId),
		ChildTId: int64(me.TId),
		RetCode:  int32(tr.RetCode),
	}
	if tr.RetFiels != nil {
		b, err := netlib.MarshalPacketNoPackId(tr.RetFiels)
		if err != nil {
			logger.Logger.Warn("ContructTxResultPacket Marshal UserData error:", err)
		} else {
			packet.CustomData = b
		}
	}
	return packet
}
