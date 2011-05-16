include $(GOROOT)/src/Make.inc

TARG=github.com/badgerodon/zmq
CGOFILES=zmq.go
CGO_CFLAGS=-I. -I"$(GOROOT)/include" -I/usr/local/include
CGO_LDFLAGS=-lzmq

include $(GOROOT)/src/Make.pkg
