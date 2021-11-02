package testcli

type Config struct {
	Server struct {
		ServiceName string   `toml:"service_name"`
		Port        int      `toml:"port"`
		Tags        []string `toml:"server_tags"`

		TCP struct {
			IdleTimeout      int `toml:"idle_timeout"`
			KeepliveInterval int `toml:"keeplive_interval"`
		} `toml:"tcp"`

		HTTP struct {
			Location    string `toml:"location"`
			LogResponse string `toml:"logResponse"`
		} `toml:"http"`

		Ratelimit []struct {
			Resource string `toml:"resource"`
			Peer     string `toml:"peer"`
			Limit    int    `toml:"limit"`
		} `toml:"ratelimit"`

		Breaker []struct {
			Resource   string `toml:"resource"`
			Open       bool   `toml:"open"`
			MinSamples int    `toml:"minsamples"`
			EThreshold int    `toml:"error_percent_threshold"`
			CThreshold int    `toml:"connsecutive_error_threshold"`
		} `toml:"breaker"`
	} `toml:"server"`

	Trace struct {
		Port    int  `toml:"port"`
		Disable bool `toml:"disable"`
	} `toml:"trace"`

	Monitor struct {
		AliveInterval int `toml:"alive_interval"`
	} `toml:"monitor"`

	Log struct {
		Level             string `toml:"level"`
		Rotate            string `toml:"rotate"`
		AccessRotate      string `toml:"access_rotate"`
		Accesslog         string `toml:"accesslog"`
		Businesslog       string `toml:"businesslog"`
		Serverlog         string `toml:"serverlog"`
		StatLog           string `toml:"statlog"`
		ErrorLog          string `toml:"errlog"`
		LogPath           string `toml:"logpath"`
		BalanceLogLevel   string `toml:"balance_log_level"`
		GenLogLevel       string `toml:"gen_log_level"`
		AccessLogOff      bool   `toml:"access_log_off"`
		BusinessLogOff    bool   `toml:"business_log_off"`
		RequestBodyLogOff bool   `toml:"request_log_off"`
		SuccStatCode      []int  `toml:"succ_stat_code"`
	} `toml:"log"`

	ServerClient        []ServerClient       `toml:"server_client"`
	KafkaConsume        []KafkaConsumeConfig `toml:"kafka_consume"`
	KafkaProducerClient []KafkaProducerItem  `toml:"kafka_producer_client"`
	Redis               []RedisConfig        `toml:"redis"`
	Database            []SQLGroupConfig     `toml:"database"`
	Circuit             []CircuitConfig      `toml:"circuit"`
	DataLog             JSONDataLogOption    `toml:"data_log"`
}
type ServerClient struct {
	APPName         *string `toml:"app_name"`
	ServiceName     string  `toml:"service_name"`
	Ipport          string  `toml:"endpoints"`
	Balancetype     string  `toml:"balancetype"`
	ProtoType       string  `toml:"proto"`
	ConnectTimeout  int     `toml:"connnect_timeout"`
	Namespace       string  `toml:"namespace"`
	ReadTimeout     int     `toml:"read_timeout"`
	WriteTimeout    int     `toml:"write_timeout"`
	MaxIdleConns    int     `toml:"max_idleconn"`
	RetryTimes      int     `toml:"retry_times"`
	SlowTime        int     `toml:"slow_time"`
	EndpointsFrom   string  `toml:"endpoints_from"`
	ConsulName      string  `toml:"consul_name"`
	LoadBalanceStat bool    `toml:"loadbalance_stat"`
	DC              string  `toml:"dc,omitempty"`
}
type JSONDataLogOption struct {
	Path     string `toml:"path"`
	Rotate   string `toml:"rotate"`
	TaskName string `toml:"task_name"`
}

type CircuitConfig struct {
	Type       string  `toml:"type"`
	Service    string  `toml:"service"`
	Resource   string  `toml:"resource"`
	End        string  `toml:"end"`
	Open       bool    `toml:"open"`
	Threshold  float64 `toml:"threshold"`
	Strategy   string  `toml:"strategy"`
	MinSamples int64   `toml:"minsamples"`
}

type KafkaProducerItem struct {
	KafkaProductConfig
	RequiredAcks string `toml:"required_acks"` // old rpc-go
	UseSync      bool   `toml:"use_sync"`      // old rpc-go
}
type KafkaProductConfig struct {
	ProducerTo     string `toml:"producer_to"`
	Broken         string `toml:"kafka_broken"`
	Topic          string `toml:"topic"`
	RetryMax       int    `toml:"retrymax"`
	RequiredAcks   string `toml:"RequiredAcks"`
	GetError       bool   `toml:"get_error"`
	GetSuccess     bool   `toml:"get_success"`
	RequestTimeout int    `toml:"request_timeout"`
}
type KafkaConsumeConfig struct {
	ConsumeFrom    string `toml:"consume_from"`
	Zookeeperhost  string `toml:"zkpoints"`
	Topic          string `toml:"topic"`
	Group          string `toml:"group"`
	Initoffset     int    `toml:"initoffset"`
	ProcessTimeout int    `toml:"process_timeout"`
	CommitInterval int    `toml:"commit_interval"`
	GetError       bool   `toml:"get_error"`
	TraceEnable    bool   `toml:"trace_enable"`
	ConsumeAll     bool   `toml:"consume_all"`
}

type SQLGroupConfig struct {
	Name      string   `toml:"name"`
	Master    string   `toml:"master"`
	Slaves    []string `toml:"slaves"`
	StatLevel string   `toml:"stat_level"`
	LogFormat string   `toml:"log_format"`
}
type RedisConfig struct {
	ServerName     string `toml:"server_name"`
	Addr           string `toml:"addr"`
	Password       string `toml:"password"`
	MaxIdle        int    `toml:"max_idle"`
	MaxActive      int    `toml:"max_active"`
	IdleTimeout    int    `toml:"idle_timeout"`
	ConnectTimeout int    `toml:"connect_timeout"`
	ReadTimeout    int    `toml:"read_timeout"`
	WriteTimeout   int    `toml:"write_timeout"`
	Database       int    `toml:"database"`
	SlowTime       int    `toml:"slow_time"`
	Retry          int    `toml:"retry"`
}
