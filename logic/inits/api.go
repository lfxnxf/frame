package inits

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/ecode"
	"github.com/lfxnxf/frame/BackendPlatform/golang/kafka"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/BackendPlatform/golang/redis"
	"github.com/lfxnxf/frame/BackendPlatform/golang/sql"
	"github.com/lfxnxf/frame/logic/inits/breaker"
	"github.com/lfxnxf/frame/logic/inits/config"
	httpclient "github.com/lfxnxf/frame/logic/inits/http/client"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/logic/inits/ratelimit"
	rpcclient "github.com/lfxnxf/frame/logic/inits/rpc/client"
	"github.com/lfxnxf/frame/logic/inits/rpc/codec"
	rpcserver "github.com/lfxnxf/frame/logic/inits/rpc/server"
	dutils "github.com/lfxnxf/frame/logic/inits/utils"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/registry"
	"github.com/hashicorp/go-multierror"
)

func (i *Inits) Shutdown() error {
	// fixme:will be blocked
	// for _, client := range i.consumeClientMap {
	// 	result = multierror.Append(result, client.Close())
	// }

	result := i.shutdown()

	_ = defaultTraceCloser.Close()
	logging.Sync()

	return result
}

func (i *Inits) shutdown() error {
	var result error
	i.producerClients.Range(func(key, value interface{}) bool {
		switch value := value.(type) {
		case *kafka.KafkaClient:
			result = multierror.Append(result, value.Close())
		case *kafka.KafkaSyncClient:
			result = multierror.Append(result, value.Close())
		}
		return true
	})
	return result
}

func (i *Inits) Config() *config.Namespace {
	namespace := i.Namespace
	if n, ok := i.namespaceConfig.Load(namespace); ok {
		return n.(*config.Namespace)
	}
	prefix := getRegistryKVPath(i.localAppServiceName)
	if len(namespace) > 0 {
		prefix = path.Join(prefix, "app")
	}
	n, _ := i.namespaceConfig.LoadOrStore(
		namespace,
		config.NewNamespace(prefix).With(namespace),
	)
	return n.(*config.Namespace)
}

func (i *Inits) ConfigInstance() config.Config {
	return i.configInstance
}

func (i *Inits) File(files ...string) error {
	return i.configInstance.LoadFile(files...)

}
func (i *Inits) Remote(paths ...string) error {
	for _, p := range paths {
		if len(p) == 0 {
			continue
		}
		p = filepath.Join(getRegistryKVPath(i.localAppServiceName), p)
		err := i.configInstance.LoadPath(p, false, "toml")
		if err != nil {
			return err
		}
	}
	return nil
}

// path param is "/config", remote path is "/service_config/link/link.time.confignew/config"
func (i *Inits) RemoteKV(path string) (string, error) {
	p := filepath.Join(getRegistryKVPath(i.localAppServiceName), path)
	return config.RemoteKV(p)
}

func (i *Inits) WatchKV(path string) config.Watcher {
	prefix := getRegistryKVPath(i.localAppServiceName)
	p := filepath.Join(prefix, path)
	return config.WatchKV(p, prefix)
}

func (i *Inits) WatchPrefix(path string) config.Watcher {
	p := filepath.Join(getRegistryKVPath(i.localAppServiceName), path)
	return config.WatchPrefix(p)
}

func (i *Inits) InjectServerClient(sc ServerClient) {
	if atomic.LoadInt32(&i.pendingServerClientTaskDone) == 0 {
		i.pendingServerClientLock.Lock()
		defer i.pendingServerClientLock.Unlock()
		if atomic.LoadInt32(&i.pendingServerClientTaskDone) == 0 {
			i.pendingServerClientTask = append(i.pendingServerClientTask, sc)
			return
		}
	}
	i.injectServerClient(sc)
}

// 对于以下方法中的service参数说明:
// 如果对应的server_client配置了app_name选项,则需要调用方保证service参数带上app_name前缀
// 如果没有配置,则保持原有逻辑,	service参数不用改动
func (i *Inits) FindServerClient(service string) (ServerClient, error) {
	if value, ok := i.serverClientMap.Load(service); ok {
		sc := value.(ServerClient)
		return sc, nil
	}
	return ServerClient{}, fmt.Errorf("client config for %s not exist", service)
}

func (i *Inits) ServiceClientWithApp(appName, serviceName string) (ServerClient, error) {
	appServiceName := dutils.MakeAppServiceName(appName, serviceName)
	return i.FindServerClient(appServiceName)
}

// RPC create a new rpc client instance, default use http protocol.
func (i *Inits) RPCFactory(name string) rpcclient.Factory {
	if c, ok := i.rpcClientMap.Load(name); ok {
		return c.(rpcclient.Factory)
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	if c, ok := i.rpcClientMap.Load(name); ok {
		return c.(rpcclient.Factory)
	}

	sc, err := i.FindServerClient(name)
	if err != nil {
		fmt.Printf("namespace %s rpcclient %s not exist, err %v\n", i.Namespace, name, err)
		logging.GenLogf("namespace %s rpcclient %s not exist, err %v", i.Namespace, name, err)
		return nil
	}
	sName := sc.ServiceName
	if sc.APPName != nil {
		if len(*sc.APPName) > 0 && *sc.APPName != _inkeApp {
			sName = fmt.Sprintf("%s.%s", *sc.APPName, sc.ServiceName)
		}
	}

	var clusterName string
	if sc.APPName != nil {
		clusterName = fmt.Sprintf("%s-http", sName)
	} else {
		clusterName = fmt.Sprintf("%s-http", dutils.MakeAppServiceName(i.App, sc.ServiceName))
	}

	client := rpcclient.HFactory(
		rpcclient.Cluster(i.Clusters.Cluster(clusterName)),
		rpcclient.Kit(DefaultKit),
		rpcclient.Tracer(defaultTracer),
		rpcclient.Codec(codec.NewJSONCodec()),
		rpcclient.MaxIdleConns(sc.MaxIdleConns),
		rpcclient.MaxIdleConnsPerHost(sc.MaxIdleConnsPerHost),
		rpcclient.DialTimeout(time.Duration(sc.ConnectTimeout)*time.Millisecond),
		rpcclient.Retries(sc.RetryTimes),
		rpcclient.RequestTimeout(time.Duration(sc.ReadTimeout)*time.Millisecond),
		rpcclient.Slow(time.Duration(sc.SlowTime)*time.Millisecond),
		rpcclient.SDName(i.localAppServiceName), // 本地服务发现名
		rpcclient.Name(sName),                   // 下游服务发现名
		rpcclient.Namespace(sc.Namespace),       // 下游所属namespace
		rpcclient.Limiter(ratelimit.NewConfig(getClientLimiterConfig(sc.Namespace, sc))),
		rpcclient.Breaker(breaker.NewConfig(getClientBreakerConfig(sc.Namespace, sc))),
	)
	i.rpcClientMap.Store(name, client)
	return client
}

func (i *Inits) RPCServer() rpcserver.Server {
	port := i.config.Server.Port
	if port == 0 {
		panic("server port is 0")
	}

	server := rpcserver.BothServer(
		i.localAppServiceName,
		port,
		rpcserver.Name(i.localAppServiceName),
		rpcserver.Tracer(defaultTracer),
		rpcserver.LoggerKit(DefaultKit),
		rpcserver.Tags(getServiceTags(i.config.Server.Tags)),
		rpcserver.Manager(i.Manager),
		rpcserver.Registry(registry.Default),
		rpcserver.Limiter(defaultServerLimiter),
		rpcserver.Breaker(defaultServerBreaker),
	)
	return server
}

func (i *Inits) HTTPClient(name string) httpclient.Client {
	if len(name) == 0 {
		if c, ok := i.httpClientMap.Load("default"); ok {
			return c.(httpclient.Client)
		}
		i.mu.Lock()
		defer i.mu.Unlock()

		if c, ok := i.httpClientMap.Load("default"); ok {
			return c.(httpclient.Client)
		}
		c := httpclient.NewClient(
			httpclient.Tracer(defaultTracer),
			httpclient.Logger(DefaultKit),
			httpclient.LocalName(i.localAppServiceName),
		)
		i.httpClientMap.Store("default", c)
		return c
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	sc, err := i.FindServerClient(name)
	if err != nil {
		fmt.Printf("namespace %s httpclient %s not exist, err %v\n", i.Namespace, name, err)
		logging.GenLogf("namespace %s httpclient %s not exist, err %v", i.Namespace, name, err)
		return nil
	}
	if sc.ProtoType == "" || sc.ProtoType == "rpc" {
		sc.ProtoType = "http"
	}
	sName := sc.ServiceName
	if sc.APPName != nil {
		if len(*sc.APPName) > 0 && *sc.APPName != _inkeApp {
			sName = fmt.Sprintf("%s.%s", *sc.APPName, sc.ServiceName)
		}
	}
	if v, ok := i.httpClientMap.Load(sName); ok {
		return v.(httpclient.Client)
	}
	var clusterName string
	if sc.APPName != nil {
		clusterName = fmt.Sprintf("%s-%s", sName, sc.ProtoType)
	} else {
		clusterName = fmt.Sprintf("%s-%s", dutils.MakeAppServiceName(i.App, sc.ServiceName), sc.ProtoType)
	}

	client := httpclient.NewClient(
		httpclient.Cluster(i.Clusters.Cluster(clusterName)),
		httpclient.Logger(DefaultKit),
		httpclient.Tracer(defaultTracer),
		httpclient.MaxIdleConns(sc.MaxIdleConns),
		httpclient.MaxIdleConnsPerHost(sc.MaxIdleConnsPerHost),
		httpclient.RetryTimes(sc.RetryTimes),
		httpclient.DialTimeout(time.Duration(sc.ConnectTimeout)*time.Millisecond),
		httpclient.RequestTimeout(time.Duration(sc.ReadTimeout)*time.Millisecond),
		httpclient.SlowTimeout(time.Duration(sc.SlowTime)*time.Millisecond),
		httpclient.KeepAliveTimeout(time.Duration(sc.KeepaliveTimeout)*time.Millisecond),
		httpclient.LocalName(i.localAppServiceName), // 本地服务发现名
		httpclient.ServiceName(sName),               // 下游服务发现名
		httpclient.Namespace(sc.Namespace),          // 下游所属namespace
		httpclient.ProtoType(sc.ProtoType),
		httpclient.Limiter(ratelimit.NewConfig(getClientLimiterConfig(sc.Namespace, sc))),
		httpclient.Breaker(breaker.NewConfig(getClientBreakerConfig(sc.Namespace, sc))),
	)

	i.httpClientMap.Store(sName, client)
	return client
}

func (i *Inits) HTTPServer() httpserver.Server {
	httpServer := httpserver.NewServer(
		httpserver.Name(i.localAppServiceName),
		httpserver.Port(i.config.Server.Port),
		httpserver.Tracer(defaultTracer),
		httpserver.Logger(DefaultKit),
		httpserver.Tags(getServiceTags(i.config.Server.Tags)),
		httpserver.Manager(i.Manager),
		httpserver.Registry(registry.Default),
		httpserver.Limiter(defaultServerLimiter),
		httpserver.Breaker(defaultServerBreaker),
		httpserver.RequestBodyLogOff(i.config.Log.RequestBodyLogOff),
		httpserver.RespBodyLogMaxSize(i.config.Log.RespBodyLogMaxSize),
		httpserver.RecoverPanic(i.config.Server.RecoverPanic),
		httpserver.ReadTimeout(time.Duration(i.config.Server.HTTP.ReadTimeout)*time.Second),
		httpserver.WriteTimeout(time.Duration(i.config.Server.HTTP.WriteTimeout)*time.Second),
		httpserver.IdleTimeout(time.Duration(i.config.Server.HTTP.IdleTimeout)*time.Second),
	)
	// default export API
	{
		// Register a router to get pprof port of this application
		httpServer.ANY(_pprofURI, func(c *httpserver.Context) {
			c.JSON(map[string]interface{}{"port": i.config.Trace.Port}, nil)
		})
		httpServer.POST("/debug/set", func(c *httpserver.Context) {
			var r struct {
				LogLevel string `json:"log_level"`
			}
			buf, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(nil, ecode.ServerErr)
				return
			}
			err = json.Unmarshal(buf, &r)
			if err != nil {
				c.JSON(nil, ecode.ServerErr)
				return
			}
			logging.SetLevelByString(r.LogLevel)
			c.JSON(nil, ecode.OK)
		})
	}
	return httpServer
}

func (i *Inits) RedisClient(name string) *redis.Redis {
	if client, ok := i.redisClients.Load(name); ok {
		if v, ok1 := client.(*redis.Redis); ok1 {
			return v
		}
	}
	fmt.Printf("namespace %s redis client for %s not exist\n", i.Namespace, name)
	logging.GenLogf("namespace %s redis client for %s not exist", i.Namespace, name)
	return nil
}

func (i *Inits) SQLClient(name string) *sql.Group {
	if client, ok := i.mysqlClients.Load(name); ok {
		if v, ok1 := client.(*sql.Group); ok1 {
			return v
		}
	}
	fmt.Printf("namespace %s mysql client for %s not exist\n", i.Namespace, name)
	logging.GenLogf("namespace %s mysql client for %s not exist", i.Namespace, name)
	return nil
}

func (i *Inits) KafkaConsumeClient(consumeFrom string) *kafka.KafkaConsumeClient {
	if client, ok := i.consumeClients.Load(consumeFrom); ok {
		if v, ok1 := client.(*kafka.KafkaConsumeClient); ok1 {
			return v
		}
	}
	fmt.Printf("namespace %s kafka consume client %s not exist\n", i.Namespace, consumeFrom)
	logging.GenLogf("namespace %s kafka consume client %s not exist", i.Namespace, consumeFrom)
	return nil
}

func (i *Inits) KafkaProducerClient(producerTo string) *kafka.KafkaClient {
	if client, ok := i.producerClients.Load(producerTo); ok {
		if v, ok := client.(*kafka.KafkaClient); ok {
			return v
		}
		fmt.Printf("namespace %s kafka producer %s type not match, should use SyncProducerClient()\n", i.Namespace, producerTo)
		logging.GenLogf("namespace %s kafka producer %s type not match, should use SyncProducerClient()", i.Namespace, producerTo)
		return nil
	}
	fmt.Printf("namespace %s kafka producer client %s to not exist\n", i.Namespace, producerTo)
	logging.GenLogf("namespace %s kafka producer client %s to not exist", i.Namespace, producerTo)
	return nil
}

func (i *Inits) SyncProducerClient(producerTo string) *kafka.KafkaSyncClient {
	if client, ok := i.producerClients.Load(producerTo); ok {
		if v, ok := client.(*kafka.KafkaSyncClient); ok {
			return v
		}
		fmt.Printf("namespace %s kafka sync producer %s type not match, should use KafkaProducerClient()\n", i.Namespace, producerTo)
		logging.GenLogf("namespace %s kafka sync producer %s type not match, should use KafkaProducerClient()", i.Namespace, producerTo)
		return nil
	}
	fmt.Printf("namespace %s kafka sync producer client %s not exist\n", i.Namespace, producerTo)
	logging.GenLogf("namespace %s kafka sync producer client %s not exist", i.Namespace, producerTo)
	return nil
}

func (i *Inits) InitKafkaProducer(kpcList []kafka.KafkaProductConfig) error {
	for _, item := range kpcList {
		if _, ok := i.producerClients.Load(item.ProducerTo); ok {
			continue
		}
		if item.UseSync {
			client, err := kafka.NewSyncProducterClient(item)
			if err != nil {
				return err
			}
			// 忽略已存在的记录
			i.producerClients.LoadOrStore(item.ProducerTo, client)
		} else {
			client, err := kafka.NewKafkaClient(item)
			if err != nil {
				return err
			}
			i.producerClients.LoadOrStore(item.ProducerTo, client)
		}
	}
	return nil
}

func (i *Inits) InitKafkaConsume(kccList []kafka.KafkaConsumeConfig) error {
	for _, item := range kccList {
		if _, ok := i.consumeClients.Load(item.ConsumeFrom); ok {
			continue
		}
		client, err := kafka.NewKafkaConsumeClient(item)
		if err != nil {
			return err
		}
		i.consumeClients.LoadOrStore(item.ConsumeFrom, client)
	}
	return nil
}

func (i *Inits) InitRedisClient(rcList []redis.RedisConfig) error {
	for _, c := range rcList {
		if _, ok := i.redisClients.Load(c.ServerName); ok {
			continue
		}
		cc := c
		client, err := redis.NewRedis(&cc)
		if err != nil {
			return err
		}
		i.redisClients.LoadOrStore(cc.ServerName, client)
	}
	return nil
}

func (i *Inits) InitSqlClient(sqlList []sql.SQLGroupConfig) error {
	for _, c := range sqlList {
		if _, ok := i.mysqlClients.Load(c.Name); ok {
			continue
		}
		g, err := sql.NewGroup(c)
		if err != nil {
			return err
		}
		_ = sql.SQLGroupManager.Add(c.Name, g)
		i.mysqlClients.LoadOrStore(c.Name, g)
	}
	return nil
}

func (i *Inits) AddSqlClient(name string, client *sql.Group) error {
	i.mysqlClients.LoadOrStore(name, client)
	return nil
}

func (i *Inits) AddRedisClient(name string, client *redis.Redis) error {
	i.redisClients.LoadOrStore(name, client)
	return nil
}

func (i *Inits) AddSyncKafkaClient(name string, client *kafka.KafkaSyncClient) error {
	i.producerClients.LoadOrStore(name, client)
	return nil
}

func (i *Inits) AddAsyncKafkaClient(name string, client *kafka.KafkaClient) error {
	i.producerClients.LoadOrStore(name, client)
	return nil
}

func (i *Inits) AddHTTPClient(name string, client httpclient.Client) error {
	i.httpClientMap.LoadOrStore(name, client)
	i.serverClientMap.LoadOrStore(name, ServerClient{ServiceName: name})
	return nil
}

func (i *Inits) initBreaker() {
	defaultServerBreaker.AddConfig(getServerBreakerConfig(i.Namespace, i.config))
	breaker.InitDefaultConfig(getDefaultBreakerConfig(i.Namespace, i.config.Server.DefaultCircuit))
}

func (i *Inits) initLimiter() {
	defaultServerLimiter.AddConfig(getServerLimiterConfig(i.Namespace, i.config))
	ratelimit.InitDefaultConfig(getDefaultLimiterConfig(i.Namespace, i.config.Server.DefaultCircuit))
}
