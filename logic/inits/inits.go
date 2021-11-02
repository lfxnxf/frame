package inits

import (
	"fmt"
	_ "net/http/pprof" // #nosec
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits/config"
	"github.com/lfxnxf/frame/logic/inits/utils"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/registry"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/upstream"
)

type Inits struct {
	// Name is discovery name, it is from deploy platform by default.
	// Name will be used to register to discovery service
	Name                        string
	Namespace                   string
	JobName                     string
	App                         string
	Version                     string
	Deps                        string
	LogDir                      string
	LogLevel                    string
	LogRotate                   string
	ConfigPath                  string
	ConfigMemory                []byte
	Clusters                    *upstream.ClusterManager
	Manager                     *registry.ServiceManager
	config                      initsConfig
	configInstance              config.Config
	redisClients                sync.Map
	mysqlClients                sync.Map
	consumeClients              sync.Map
	producerClients             sync.Map
	serverClientMap             sync.Map
	httpClientMap               sync.Map
	namespaceConfig             sync.Map
	rpcClientMap                sync.Map
	mu                          sync.Mutex
	localAppServiceName         string
	pendingServerClientTask     []ServerClient
	pendingServerClientLock     sync.Mutex
	pendingServerClientTaskDone int32
	initOnce                    sync.Once
	namespaceDir                string
}

func New() *Inits {
	return &Inits{
		Name:                        "",
		Namespace:                   "",
		App:                         "",
		Version:                     "",
		LogDir:                      "logs",
		LogLevel:                    "debug",
		LogRotate:                   "hour",
		ConfigPath:                  "",
		Clusters:                    upstream.NewClusterManager(),
		Manager:                     nil,
		configInstance:              nil,
		pendingServerClientTask:     nil,
		pendingServerClientLock:     sync.Mutex{},
		pendingServerClientTaskDone: 0,
		initOnce:                    sync.Once{},
	}
}

func (i *Inits) Init(options ...Option) {
	i.initOnce.Do(func() {
		for _, opt := range options {
			opt(i)
		}
		if len(i.JobName) == 0 {
			i.JobName = strings.TrimSpace(readFile(".jobname"))
		}
		if len(i.Deps) == 0 {
			i.Deps = readFile(".deps")
		}
		if len(i.App) == 0 {
			i.App = strings.TrimSpace(readFile(".app"))
		}
		if len(i.Version) == 0 {
			i.Version = strings.TrimSpace(readFile(".version"))
		}

		// 读取本地配置
		var cc []byte
		if len(i.ConfigMemory) > 0 {
			cc = i.ConfigMemory
		} else {
			cc = i.loadLocalConfig()
		}

		// 设置服务发现名
		i.initLocalAppServiceName(cc)

		// 处理远程开关逻辑
		initRemoteFirst(i.localAppServiceName)

		i.pendingServerClientLock.Lock()
		pending := i.pendingServerClientTask
		i.pendingServerClientLock.Unlock()

		if len(cc) > 0 {
			if err := i.initConfigInstance(); err != nil {
				panic(err)
			}

			// logger,consul backend,tracer都只初始化一次
			if len(i.ConfigPath) > 0 {
				i.initDefaultOnce()
			}

			i.Manager = registry.NewServiceManager(logging.Log(logging.BalanceLoggerName))

			// init middleware client
			if err := i.initMiddleware(); err != nil {
				panic(err)
			}

			// inject service client from config
			i.pendingServerClientLock.Lock()
			pending = append(pending, i.config.ServerClient...)
			i.pendingServerClientTask = nil
			atomic.StoreInt32(&i.pendingServerClientTaskDone, 1)
			i.pendingServerClientLock.Unlock()
		}

		for _, sc := range pending {
			i.injectServerClient(sc)
		}

		curTime := time.Now().Format(utils.TimeFormat)
		fmt.Printf("%s init inits success app:%s name:%s namespace:%s config:%s\n",
			curTime, i.App, i.Name, i.Namespace, i.ConfigPath)
	})
}
