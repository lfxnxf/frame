package _testing

import (
	"fmt"
	"github.com/lfxnxf/frame/logic/inits/rpc/_testing/newpb"
	"github.com/lfxnxf/frame/logic/inits/rpc/client"
	"github.com/lfxnxf/frame/logic/inits/rpc/codec"
	"github.com/lfxnxf/frame/logic/inits/rpc/server"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"sync"
	"testing"
	"time"

	config "github.com/lfxnxf/frame/tpc/inf/go-upstream/config"
	upstream "github.com/lfxnxf/frame/tpc/inf/go-upstream/upstream"
)

func TestBinaryCS(t *testing.T) {
	assert := assert.New(t)

	server := NNew(server.BinaryServer(
		server.Address("0.0.0.0:10000"),
		server.Codec(codec.NewProtoCodec()),
	))

	clusterName := "test_client"
	config := config.NewCluster()
	config.Name = clusterName
	config.StaticEndpoints = "127.0.0.1:10000"

	manager := upstream.NewClusterManager()
	assert.Nil(manager.InitService(config))

	client := client.BinaryClient(
		client.Cluster(manager.Cluster(clusterName)),
		client.Codec(codec.NewProtoCodec()),
	)

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Millisecond * 100)

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			m := fmt.Sprintf("%d", i)
			response := &newpb.EchoResponse{}
			request := &newpb.EchoRequest{
				Message: proto.String(m),
			}
			err := client.Invoke(context.TODO(), "echo.EchoService.Echo", request, response)
			assert.Nil(err)
			assert.EqualValues(m, response.GetResponse())
		}()
	}
	wg.Wait()
}

func TestHTTPCS(t *testing.T) {
	assert := assert.New(t)

	server := NNew(server.HTTPServer(
		server.Address("0.0.0.0:10001"),
		server.Codec(codec.NewProtoCodec()),
	))

	clusterName := "test_client"
	config := config.NewCluster()
	config.Name = clusterName
	config.StaticEndpoints = "127.0.0.1:10001"

	manager := upstream.NewClusterManager()
	assert.Nil(manager.InitService(config))

	client := client.HTTPClient(
		client.Cluster(manager.Cluster(clusterName)),
		client.Codec(codec.NewProtoCodec()),
	)

	go func() {
		if err := server.Start(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Millisecond * 100)

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			m := fmt.Sprintf("%d", i)
			response := &newpb.EchoResponse{}
			request := &newpb.EchoRequest{
				Message: proto.String(m),
			}
			err := client.Invoke(context.TODO(), "echo.EchoService.Echo", request, response)
			assert.Nil(err)
			assert.EqualValues(m, response.GetResponse())
		}()
	}
	wg.Wait()

}
