package netlib

import (
	"errors"
	"github.com/breezedup/goserver/core/builtin/protocol"
	"github.com/breezedup/goserver/core/logger"
)

// ServerError represents an error that has been returned from
// the remote side of the RPC connection.
type ServerError string

func (e ServerError) Error() string {
	return string(e)
}

var ErrShutdown = errors.New("connection is shut down")
var ErrSendBufFull = errors.New("sendbuf is full")
var ErrUnsupportRpc = errors.New("only inner session support rpc")

// If set, print log statements for internal and I/O errors.
var debugRPCLog = false

// Call represents an active RPC.
type Call struct {
	ServiceMethod string      // The name of the service and method to call.
	Args          interface{} // The argument to the function (*struct).
	Reply         interface{} // The reply from the function (*struct).
	Error         error       // After completion, the error status.
	Done          chan *Call  // Receives *Call when Go is complete.
}

func (call *Call) done() {
	select {
	case call.Done <- call:
		// ok
	default:
		// We don't want to block here. It is the caller's responsibility to make
		// sure the channel has enough buffer space. See comment in Go().
		if debugRPCLog {
			logger.Logger.Debugf("rpc: discarding Call reply due to insufficient Done chan capacity")
		}
	}
}

// Go invokes the function asynchronously. It returns the Call structure representing
// the invocation. The done channel will signal when the call is complete by returning
// the same Call object. If done is nil, Go will allocate a new channel.
// If non-nil, done must be buffered or Go will deliberately crash.
func (s *Session) GoRpc(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
	call := new(Call)
	call.ServiceMethod = serviceMethod
	call.Args = args
	call.Reply = reply
	if !s.sc.IsInnerLink {
		call.Error = ErrUnsupportRpc
		return call
	}
	if done == nil {
		done = make(chan *Call, 1) // buffered.
	} else {
		// If caller passes done != nil, it must arrange that
		// done has enough buffer for the number of simultaneous
		// RPCs that will be using that channel. If the channel
		// is totally unbuffered, it's best not to run at all.
		if cap(done) == 0 {
			logger.Logger.Criticalf("rpc: done channel is unbuffered")
		}
	}
	call.Done = done
	s.sendRpcReq(call)
	return call
}

// Call invokes the named function, waits for it to complete, and returns its error status.
func (s *Session) CallRpc(serviceMethod string, args interface{}, reply interface{}) error {
	call := <-s.GoRpc(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}

func (s *Session) sendRpcReq(call *Call) {
	// Register this call.
	s.mutex.Lock()
	if !s.isConned || s.quit || s.shutSend || s.shutRecv {
		s.mutex.Unlock()
		call.Error = ErrShutdown
		call.done()
		return
	}
	seq := s.seq
	s.seq++
	s.pending[seq] = call
	s.mutex.Unlock()

	// Encode and send the request.
	req := &protocol.RpcRequest{
		ServiceMethod: call.ServiceMethod,
		Seq:           seq,
	}
	req.Args, _ = Gob.Marshal(call.Args)

	if !s.Send(int(protocol.CoreBuiltinPacketID_PACKET_SS_RPC_REQ), req, true) {
		s.mutex.Lock()
		call = s.pending[seq]
		delete(s.pending, seq)
		s.mutex.Unlock()
		if call != nil {
			call.Error = ErrSendBufFull
			call.done()
		}
	}
}

func (s *Session) onRpcResp(resp *protocol.RpcResponse) {
	s.mutex.Lock()
	call := s.pending[resp.Seq]
	delete(s.pending, resp.Seq)
	s.mutex.Unlock()

	switch {
	case call == nil:
		// We've got no pending call. That usually means that
		// WriteRequest partially failed, and call was already
		// removed; response is a server telling us about an
		// error reading request body. We should still attempt
		// to read error body, but there's no one to give it to.
	case resp.Error != "":
		// We've got an error response. Give this to the request;
		// any subsequent requests will get the ReadResponseBody
		// error if there is one.
		call.Error = ServerError(resp.Error)
		call.done()
	default:
		call.Error = Gob.Unmarshal(resp.Reply, call.Reply)
		call.done()
	}
}

func init() {
	RegisterHandler(int(protocol.CoreBuiltinPacketID_PACKET_SS_RPC_RESP), HandlerWrapper(func(s *Session, packetid int, data interface{}) error {
		if resp, ok := data.(*protocol.RpcResponse); ok {
			s.onRpcResp(resp)
		}
		return nil
	}))
	RegisterFactory(int(protocol.CoreBuiltinPacketID_PACKET_SS_RPC_RESP), PacketFactoryWrapper(func() interface{} {
		return &protocol.RpcResponse{}
	}))
}
