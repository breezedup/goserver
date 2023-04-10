package main

import (
	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/netlib"
	"github.com/breezedup/goserver/mmo/protocol"
	"github.com/breezedup/goserver/srvlib"
)

var (
	SessionHandlerClientLoadName = "handler-client-load"
)

type SessionHandlerClientLoad struct {
	netlib.BasicSessionHandler
}

func (sfcl SessionHandlerClientLoad) GetName() string {
	return SessionHandlerClientLoadName
}

func (sfcl *SessionHandlerClientLoad) GetInterestOps() uint {
	return 1<<netlib.InterestOps_Opened | 1<<netlib.InterestOps_Closed
}

func (sfcl *SessionHandlerClientLoad) OnSessionOpened(s *netlib.Session) {
	sfcl.reportLoad(s)
}

func (sfcl *SessionHandlerClientLoad) OnSessionClosed(s *netlib.Session) {
	sfcl.reportLoad(s)

}

func (sfcl *SessionHandlerClientLoad) reportLoad(s *netlib.Session) {
	sc := s.GetSessionConfig()
	pack := &protocol.ServerLoad{
		SrvType: int32(sc.Type),
		SrvId:   int32(sc.Id),
		CurLoad: int32(srvlib.ClientSessionMgrSington.Count()),
	}
	srvlib.ServerSessionMgrSington.Broadcast(int(protocol.MmoPacketID_PACKET_SC_GATEINFO), pack, netlib.Config.SrvInfo.AreaID, srvlib.BalanceServerType)
	logger.Logger.Tracef("SessionHandlerClientLoad.reportLoad %v", pack)
}

func init() {
	netlib.RegisteSessionHandlerCreator(SessionHandlerClientLoadName, func() netlib.SessionHandler {
		return &SessionHandlerClientLoad{}
	})
}
