package ikio

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestLineServer(t *testing.T) {
	server := NewServer(
		CustomCodecOption(func() Codec { return &LineCodec{} }),
	)
	ln, err := net.Listen("tcp4", "127.0.0.1:12345")
	if err != nil {
		t.Fatalf("listen error %s", err)
	}
	server.Register(0, func(ctx context.Context, conn WriteCloser) {
		message := MessageFromContext(ctx).(*LinePacket)
		conn.Write(message)
		conn.(*ServerChannel).RunAfter(1*time.Second, 0, func(time.Time, WriteCloser) {
			conn.(*ServerChannel).Close()
		})
	}, HandlePooledRandom)
	go func() {
		conn, err := net.Dial("tcp4", "127.0.0.1:12345")
		if err != nil {
			t.Fatalf("dial server error %s\n", err)
		}
		channel := NewClientChannel(conn,
			CustomCodecOption(func() Codec { return &LineCodec{} }),
			OnConnectOption(func(conn WriteCloser) bool {
				fmt.Printf("connect success\n")
				return true
			}), OnMessageOption(func(p Packet, cc WriteCloser) {
				message := p.(*LinePacket).Payload
				fmt.Printf("recv message %q\n", message)
			}), OnCloseOption(func(cc WriteCloser) {
				fmt.Printf("connect close\n")
			}),
		)
		channel.Start()
		channel.RunAfter(100*time.Millisecond, 100*time.Millisecond, func(t time.Time, cc WriteCloser) {
			channel.Write(&LinePacket{Payload: []byte("hello world")})
		})
	}()
	err = server.Start(ln)
	fmt.Printf("server start %s, err %s", ln.Addr(), err)
}

func TestLinkServer(t *testing.T) {
	server := NewServer(
		CustomCodecOption(func() Codec { return &LinkCodec{} }),
	)
	ln, err := net.Listen("tcp4", "127.0.0.1:22345")
	if err != nil {
		t.Fatalf("listen error %s", err)
	}
	server.Register(0, func(ctx context.Context, conn WriteCloser) {
		message := MessageFromContext(ctx).(*LinkPacket)
		conn.Write(message)
		// conn.(*ServerChannel).RunAfter(1*time.Second, 0, func(time.Time, WriteCloser) {
		// 	conn.(*ServerChannel).Close()
		// })
	}, HandlePooledRandom)
	go func() {
		conn, err := net.Dial("tcp4", "127.0.0.1:22345")
		if err != nil {
			t.Fatalf("dial server error %s\n", err)
		}
		channel := NewClientChannel(conn,
			CustomCodecOption(func() Codec { return &LinkCodec{} }),
			OnConnectOption(func(conn WriteCloser) bool {
				fmt.Printf("connect success\n")
				return true
			}), OnMessageOption(func(p Packet, cc WriteCloser) {
				message := p.(*LinkPacket).Body
				fmt.Printf("recv message %q\n", message)
			}), OnCloseOption(func(cc WriteCloser) {
				fmt.Printf("on connect close\n")
			}),
		)
		channel.Start()
		channel.RunAfter(100*time.Millisecond, 100*time.Millisecond, func(t time.Time, cc WriteCloser) {
			channel.Write(&LinkPacket{Body: []byte("hello world"), Cmd: 0})
		})
	}()
	err = server.Start(ln)
	fmt.Printf("server start %s, err %s", ln.Addr(), err)
}

func TestRPCServer(t *testing.T) {
	server := NewServer(
		CustomCodecOption(func() Codec { return &RPCCodec{} }),
	)
	ln, err := net.Listen("tcp4", "127.0.0.1:22345")
	if err != nil {
		t.Fatalf("listen error %s", err)
	}
	server.Register(0, func(ctx context.Context, conn WriteCloser) {
		message := MessageFromContext(ctx).(*RPCPacket)
		// fmt.Printf("recv 0 packet %v\n", message)
		// conn.Write(message)
		ServerFromContext(ctx).Broadcast(message, func(int64) bool { return false } )
		// conn.(*ServerChannel).RunAfter(1*time.Second, 0, func(time.Time, WriteCloser) {
		// 	conn.(*ServerChannel).Close()
		// })
	}, HandlePooledRandom)
	server.Register(RPCNegoMessageType, func(ctx context.Context, conn WriteCloser) {
		message := MessageFromContext(ctx).(*RPCNegoPacket)
		fmt.Printf("server recv nego packet %v\n", message)
		conn.Write(message)
	}, HandlePooledRandom)
	sig := make(chan os.Signal, 0)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		server.Stop()
	}()
	for i := 0; i < 1; i++ {
		go func(i int) {
			conn, err := net.Dial("tcp4", "127.0.0.1:22345")
			if err != nil {
				t.Fatalf("dial server error %s\n", err)
			}
			channel := NewClientChannel(conn,
				CustomCodecOption(func() Codec { return &RPCCodec{} }),
				OnConnectOption(func(conn WriteCloser) bool {
					fmt.Printf("%d connect success\n", i)
					return true
				}), OnMessageOption(func(p Packet, cc WriteCloser) {
					// message := p.(*RPCPacket).Payload
					switch p.(type) {
					case *RPCPacket:
						fmt.Printf("client #%d recv packet message %q\n", i, p.(*RPCPacket).Payload)
					case *RPCNegoPacket:
						fmt.Printf("client #%d recv nego message %d\n", i, p.(*RPCNegoPacket).Flag)
					}
					// fmt.Printf("client recv message %v\n", p)
				}), OnCloseOption(func(cc WriteCloser) {
					fmt.Printf("#%d on connect close\n", i)
				}),
			)
			channel.Start()
			channel.RunAfter(100*time.Millisecond, 0, func(t time.Time, cc WriteCloser) {
				// channel.Write(&RPCPacket{Payload: []byte("hello world"), Tp: 0})
				channel.Write(&RPCNegoPacket{Magic: 1234, ID: []byte("Test")})
			})
			timer := channel.RunAfter(1000*time.Millisecond, 10*time.Millisecond, func(t time.Time, cc WriteCloser) {
				//fmt.Printf("write rpc packet \n")
				channel.Write(&RPCPacket{Payload: []byte("hello world"), Tp: 0})
			})
			time.Sleep(10 * time.Second)
			channel.CancelTimer(timer)
			fmt.Printf("#%d timer close\n", i)
		}(i)
	}
	err = server.Start(ln)
	fmt.Printf("server start %s, err %s", ln.Addr(), err)
}
