package rpc

import (
	"context"
	"fmt"
	"strings"

	"github.com/lfxnxf/frame/BackendPlatform/golang/kafka"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/BackendPlatform/golang/redis"
	"github.com/lfxnxf/frame/BackendPlatform/golang/sql"
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/lfxnxf/frame/logic/inits/rpc/client"
)

var EnableTracing = true

var logdirGlobal = "logs"

// func StartMonitor(c Config) {}

func InitFrameworkUtils(c Config) {
	// nothing to do, inits init do it.
}

// SetBusinessLogFile  set client/server access log file path
func SetBusinessLogFile(path string) {
	if inits.DefaultKit.B() == nil {
		return
	}
	inits.DefaultKit.B().SetOutputPath(path)
}

func SetAccessRotateType(rotate string) {
	if inits.DefaultKit.A() == nil {
		return
	}
	if logging.DAY_ROTATE == rotate {
		inits.DefaultKit.A().SetRotateByDay()
		return
	}
	inits.DefaultKit.A().SetRotateByHour()
}

func SetAccessLogFile(path string) {
	if inits.DefaultKit.A() == nil {
		return
	}
	inits.DefaultKit.A().SetOutputPath(path)
}

func SyncLog() {
	if inits.DefaultKit.A() != nil {
		inits.DefaultKit.A().Sync()
	}
	if inits.DefaultKit.B() != nil {
		inits.DefaultKit.B().Sync()
	}
	logging.Sync()
}

func NewRemoteTomlConfig() (*ConfigToml, error) {
	return GetTomlConfig(), nil
}

func Unregister() {
	inits.Default.Manager.Deregister()
}

func SetServiceName(name string) {
	inits.Default.Name = name
}

func InitRedisClient(redisConfigs []redis.RedisConfig) error {
	return inits.Default.InitRedisClient(redisConfigs)
}

func InitKafkaConsume(consumeConfigs []kafka.KafkaConsumeConfig) error {
	return inits.Default.InitKafkaConsume(consumeConfigs)
}

func InitKafkaProducer(producerConfigs []kafka.KafkaProductConfig) error {
	return inits.Default.InitKafkaProducer(producerConfigs)
}

func InitSQLClient(sqlConfig []sql.SQLGroupConfig) error {
	return inits.Default.InitSqlClient(sqlConfig)
}

func CloseAllClient() error {
	// should be delete
	return nil
}

func InitLoadBalance(serverClientList []ServerClient) error {
	for _, sc := range serverClientList {
		inits.Default.InjectServerClient(sc)
	}
	return nil
}

func GetRedis(service string) (r *redis.Redis, err error) {
	r = inits.RedisClient(context.TODO(), service)
	if r == nil {
		return nil, fmt.Errorf("redis client for %s not exist", service)
	}
	return r, nil
}

func GetKafkaConsumeClient(consumeFrom string) (k *kafka.KafkaConsumeClient, err error) {
	k = inits.KafkaConsumeClient(context.TODO(), consumeFrom)
	if k == nil {
		return nil, fmt.Errorf("kfk consume client for %s not exist", consumeFrom)
	}
	return k, nil
}

func GetSyncProducerClient(producerTo string) (k *kafka.KafkaSyncClient, err error) {
	k = inits.SyncProducerClient(context.TODO(), producerTo)
	if k == nil {
		return nil, fmt.Errorf("kfk sync client for %s not exist", producerTo)
	}
	return k, nil
}

func GetKafkaProducerClient(producerTo string) (k *kafka.KafkaClient, err error) {
	k = inits.KafkaProducerClient(context.TODO(), producerTo)
	if k == nil {
		return nil, fmt.Errorf("kfk client for %s not exist", producerTo)
	}
	return k, nil
}

func Invoke(ctx context.Context, method string, args, reply interface{}, cc *ClientConn) error {
	if value, ok := cc.endpoints.Load(method); !ok {
		cc.mu.Lock()

		if value, ok := cc.endpoints.Load(method); ok {
			cc.mu.Unlock()
			return value.(client.Client).Invoke(ctx, args, reply)
		}

		c := cc.factory.Client(method)
		cc.endpoints.Store(method, c)

		cc.mu.Unlock()
		return c.Invoke(ctx, args, reply)
	} else {
		return value.(client.Client).Invoke(ctx, args, reply)
	}
}

func getServiceTags(tags []string) map[string]string {
	serviceTags := make(map[string]string)
	for _, t := range tags {
		kvs := strings.SplitN(t, "=", 2)
		if len(kvs) > 1 {
			serviceTags[strings.TrimSpace(kvs[0])] = strings.TrimSpace(kvs[1])
		}
	}
	return serviceTags
}
