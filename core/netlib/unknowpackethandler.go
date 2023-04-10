package netlib

var (
	unknowPacketHandlerCreatorPool = make(map[string]UnknowPacketHandlerCreator)
)

type UnknowPacketHandlerCreator func() UnknowPacketHandler

type UnknowPacketHandler interface {
	OnUnknowPacket(s *Session, packetid int, logicNo uint32, data []byte) bool //run in session receive goroutine
}

type UnknowPacketHandlerWrapper func(session *Session, packetid int, logicNo uint32, data []byte) bool

func (hw UnknowPacketHandlerWrapper) OnUnknowPacket(session *Session, packetid int, logicNo uint32, data []byte) bool {
	return hw(session, packetid, logicNo, data)
}

func RegisteUnknowPacketHandlerCreator(name string, ephc UnknowPacketHandlerCreator) {
	if ephc == nil {
		return
	}
	if _, exist := unknowPacketHandlerCreatorPool[name]; exist {
		panic("repeate registe ErrorPacketHandler:" + name)
	}

	unknowPacketHandlerCreatorPool[name] = ephc
}

func GetUnknowPacketHandlerCreator(name string) UnknowPacketHandlerCreator {
	if ephc, exist := unknowPacketHandlerCreatorPool[name]; exist {
		return ephc
	}
	return nil
}
