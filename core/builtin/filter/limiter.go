package filter

import (
	"github.com/breezedup/goserver/core/netlib"
	"golang.org/x/time/rate"
)

var (
	LimiterFilterName       = "session-filter-limiter"
	SessionAttributeLimiter = &LimiterFilter{}
)

type LimiterFilter struct {
	netlib.BasicSessionFilter
}

func (lf *LimiterFilter) GetName() string {
	return LimiterFilterName
}

func (lf *LimiterFilter) GetInterestOps() uint {
	return 1 << netlib.InterestOps_Received
}

func (lf *LimiterFilter) OnPacketReceived(s *netlib.Session, packetid int, logicNo uint32, packet interface{}) bool {
	attr := s.GetAttribute(SessionAttributeLimiter)
	if attr != nil {
		if pool, ok := attr.(map[int]*rate.Limiter); ok && pool != nil {
			if limiter, exist := pool[packetid]; exist && limiter != nil {
				if !limiter.Allow() {
					return false
				}
			}
		}
	}
	return true
}

func init() {
	netlib.RegisteSessionFilterCreator(LimiterFilterName, func() netlib.SessionFilter {
		return &LimiterFilter{}
	})
}
