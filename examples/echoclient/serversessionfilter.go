// serversessionfilter
package main

import (
	"time"

	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/netlib"
	"github.com/breezedup/goserver/examples/protocol"
)

var (
	ServerSessionFilterName = "serversessionfilter"
)

type ServerSessionFilter struct {
	netlib.BasicSessionFilter
}

func (ssf ServerSessionFilter) GetName() string {
	return ServerSessionFilterName
}

func (ssf *ServerSessionFilter) GetInterestOps() uint {
	return 1 << netlib.InterestOps_Opened
}

func (ssf *ServerSessionFilter) OnSessionOpened(s *netlib.Session) bool {
	logger.Logger.Trace("(ssf *ServerSessionFilter) OnSessionOpened")
	packet := &protocol.CSPacketPing{
		TimeStamb: time.Now().Unix(),
		Message:   []byte("=1234567890abcderghijklmnopqrstuvwxyz="),
	}
	//for i := 0; i < 1024*32; i++ {
	//	packet.Message = append(packet.Message, byte('x'))
	//}
	s.Send(int(protocol.PacketID_PACKET_CS_PING), packet)
	return true
}

func init() {
	netlib.RegisteSessionFilterCreator(ServerSessionFilterName, func() netlib.SessionFilter {
		return &ServerSessionFilter{}
	})
}
