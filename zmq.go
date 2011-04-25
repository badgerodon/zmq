package zmq

// #include <zmq.h>
// #include <stdlib.h>
import "C"

import (
	"os"
	"unsafe"
)

const (
	ENOTSUP = C.ENOTSUP
	EPROTONOSUPPORT = C.EPROTONOSUPPORT
	ENOBUFS = C.ENOBUFS
	ENETDOWN = C.ENETDOWN
	EADDRINUSE = C.EADDRINUSE
	EADDRNOTAVAIL = C.EADDRNOTAVAIL
	ECONNREFUSED = C.ECONNREFUSED
	EINPROGRESS = C.EINPROGRESS
	EMTHREAD = C.EMTHREAD
	EFSM = C.EFSM
	ENOCOMPATPROTO = C.ENOCOMPATPROTO
	ETERM = C.ETERM
	EINVAL = C.EINVAL
	
	PAIR = C.ZMQ_PAIR
	PUB = C.ZMQ_PUB
	SUB = C.ZMQ_SUB
	REQ = C.ZMQ_REQ
	REP = C.ZMQ_REP
	XREQ = C.ZMQ_XREQ
	XREP = C.ZMQ_XREP
	PULL = C.ZMQ_PULL
	PUSH = C.ZMQ_PUSH
	HWM = C.ZMQ_HWM
	SWAP = C.ZMQ_SWAP
	AFFINITY = C.ZMQ_AFFINITY
	IDENTITY = C.ZMQ_IDENTITY
	SUBSCRIBE = C.ZMQ_SUBSCRIBE
	UNSUBSCRIBE = C.ZMQ_UNSUBSCRIBE
	RATE = C.ZMQ_RATE
	RECOVERY_IVL = C.ZMQ_RECOVERY_IVL
	MCAST_LOOP = C.ZMQ_MCAST_LOOP
	SNDBUF = C.ZMQ_SNDBUF
	RCVBUF = C.ZMQ_RCVBUF
	RCVMORE = C.ZMQ_RCVMORE
	NOBLOCK = C.ZMQ_NOBLOCK
	SNDMORE = C.ZMQ_SNDMORE
	POLLIN = C.ZMQ_POLLIN
	POLLOUT = C.ZMQ_POLLOUT
	POLLERR = C.ZMQ_POLLERR
	STREAMER = C.ZMQ_STREAMER
	FORWARDER = C.ZMQ_FORWARDER
	QUEUE = C.ZMQ_QUEUE
)

type (
	ZError struct {
		n C.int
	}
	ZContext struct {
		Ptr unsafe.Pointer
	}
	ZSocket struct {
		Ptr unsafe.Pointer
	}
	ZMessage struct {
		Ptr *C.zmq_msg_t
	}
	ZMessageData struct {
		Ptr unsafe.Pointer
	}
)

func (this ZError) String() string {
	return C.GoString(C.zmq_strerror(this.n))
}
func error() os.Error {
	errnum := C.zmq_errno()
	return ZError{errnum}
}
func handle(retval C.int) os.Error {
	if retval == C.int(0) {
		return nil
	}
	return error()
}
func Bind(socket ZSocket, endpoint string) os.Error {
	return handle(C.zmq_bind(socket.Ptr, C.CString(endpoint)))
}
func Close(socket ZSocket) os.Error {
	return handle(C.zmq_close(socket.Ptr))
}
func Connect(socket ZSocket, endpoint string) os.Error {
	return handle(C.zmq_connect(socket.Ptr, C.CString(endpoint)))
}
func Device(device int, frontend ZSocket, backend ZSocket) os.Error {
	return handle(C.zmq_device(C.int(device), frontend.Ptr, backend.Ptr))
}
//func GetSocketOption(socket *Socket,
func Init(io_threads int) (ZContext, os.Error) {
	ptr := C.zmq_init(C.int(io_threads))
	if ptr == nil {
		return ZContext{nil}, error()
	}
	return ZContext{ptr}, nil
}
func MessageClose(message ZMessage) os.Error {
	return handle(C.zmq_msg_close(message.Ptr))
}
func MessageCopy(dest ZMessage, src ZMessage) os.Error {
	return handle(C.zmq_msg_copy(dest.Ptr, src.Ptr))
}
func MessageData(message ZMessage) ZMessageData {
	return ZMessageData{C.zmq_msg_data(message.Ptr)}
}
//zmq_msg_init_data
func MessageInitSize(message ZMessage, size int) os.Error {
	return handle(C.zmq_msg_init_size(message.Ptr, C.size_t(size)))
}
func MessageInit(message ZMessage) os.Error {
	return handle(C.zmq_msg_init(message.Ptr))
}
func MessageMove(dest ZMessage, src ZMessage) os.Error {
	return handle(C.zmq_msg_move(dest.Ptr, src.Ptr))
}
func MessageSize(message ZMessage) int {
	return int(C.zmq_msg_size(message.Ptr))
}
//func Poll(items []*PollItem, timeout int64) os.Error {
//	return handle(C.zmq_poll(
//}
func Receive(socket ZSocket, message ZMessage, flags int) os.Error {
	return handle(C.zmq_recv(socket.Ptr, message.Ptr, C.int(flags)))
}
func Send(socket ZSocket, message ZMessage, flags int) os.Error {
	return handle(C.zmq_send(socket.Ptr, message.Ptr, C.int(flags)))
}
// SetSocketOption
func Socket(context ZContext, socketType int) (ZSocket, os.Error) {
	ptr := C.zmq_socket(context.Ptr, C.int(socketType))
	if ptr == nil {
		return ZSocket{nil}, error()
	}
	return ZSocket{ptr}, nil
}
func Term(context ZContext) os.Error {
	return handle(C.zmq_term(context.Ptr))
}
func Version() (int, int, int) {
	var major C.int
	var minor C.int
	var patch C.int
	C.zmq_version(&major, &minor, &patch)
	return int(major), int(minor), int(patch)
}