package ikio

import (
	"testing"
	"golang.org/x/net/context"

	"net"
	"fmt"
	"log"
	"sync"
	"time"
	"io"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

func testGetConn(address string) net.Conn {
	c, err := net.Dial("tcp4", address)
	if err != nil {
		panic(err)
	}
	return c
}

func testFakeServer(port string) {
	ln, err := net.Listen("tcp4", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("%s\n", err)
			}
			go func(c net.Conn) {
				io.Copy(c, c)
				c.Close()
			}(conn)
		}
	}()
}

func TestChannel(t *testing.T) {
	testFakeServer("10000")
	conn := testGetConn("127.0.0.1:10000")
	optss := options{
		onConnect: func(wc WriteCloser) bool {
			logging.Infof("onConnect")
			return true
		},
		onClose: func(wc WriteCloser) {
			logging.Infof("onClose")
		},
		codec: func() Codec {
			return &RPCCodec{}
		},
		logger: logging.New(),
	}

	opts := channelOption{
		options: optss,
		connID: 1,
		waitGroup: &sync.WaitGroup{},
		timingWheel: nil,
	}
	channel := NewChannel(context.TODO(), conn, opts)

	opts.waitGroup.Add(1)
	channel.Start()
	go func() {
		time.Sleep(time.Second)
		channel.Close()
	}()
	opts.waitGroup.Wait()
}

type fakeConn struct {

}

func (f *fakeConn) Read(b []byte) (n int, err error) {
	return cap(b) - len(b), nil
}

func (f *fakeConn) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func (f *fakeConn) Close() error {
	return nil
}

type fakeAddr struct {}

func (a fakeAddr) String() string {
	return "fakeAddr"
}

func (a fakeAddr) Network() string {
	return "fakeAddr"
}

func (f *fakeConn) LocalAddr() net.Addr {
	return fakeAddr{}
}

func (f *fakeConn) RemoteAddr() net.Addr {
	return fakeAddr{}
}

func (f *fakeConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *fakeConn) SetWriteDeadline(t time.Time) error {
	return nil
}

type fakePacket struct {

}

func (fakePacket) Type() int32 {
	return 1
}

func (fakePacket) Serialize() ([]byte, error) {
	return nil, nil
}

func BenchmarkWrite(b *testing.B) {
	//testFakeServer("11000")
	//conn := testGetConn("127.0.0.1:11000")
	conn := &fakeConn{}
	optss := options{
		onConnect: func(wc WriteCloser) bool {
			logging.Infof("onConnect")
			return true
		},
		onClose: func(wc WriteCloser) {
			logging.Infof("onClose")
		},
		codec: func() Codec {
			return &RPCCodec{}
		},
		logger: logging.New(),
	}

	opts := channelOption{
		options: optss,
		connID: 1,
		waitGroup: &sync.WaitGroup{},
		timingWheel: nil,
	}
	channel := NewChannel(context.TODO(), conn, opts)

	opts.waitGroup.Add(1)
	channel.Start()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := channel.Write(fakePacket{})
		if err != nil && err != ErrWouldBlock {
			panic(err)
		}
	}
}
