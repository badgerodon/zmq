include $(GOROOT)/src/Make.inc

TARG=github.com/badgerodon/zmq
CGOFILES=zmq.go
CGO_CFLAGS=-I. -I"$(GOROOT)/include" -I/usr/local/include
#CGO_LDFLAGS=`locate libzmq.a -n 1` `locate libstdc++.a -n 1` `locate libuuid.a -n 1` -lpthread
CGO_LDFLAGS=-lzmq

include $(GOROOT)/src/Make.pkg
