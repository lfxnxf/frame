package rpc

import (
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/kafka"
	"github.com/lfxnxf/frame/BackendPlatform/golang/redis"
	"github.com/lfxnxf/frame/BackendPlatform/golang/sql"
	"github.com/lfxnxf/frame/logic/inits"
)

type Config interface {
	Port() int
	String(key string) (string, bool)
	StringWithDefault(key, defaultVal string) string

	Int(key string) (int, bool)

	NegotiateTimeout() time.Duration

	IntWithDefault(key string, value int) int

	ServerLogPath() string
	ServerLogLevel() string
	BusinessLogPath() string
	AccessLogPath() string
	Rotate() string
	AccessRotate() string
	LogPath() string

	StatLogPath() string
	StatMetricName() string

	IdleTimeout() time.Duration
	KeepAliveInterval() time.Duration

	TracePort() int

	HTTPServeLocation() string
	HTTPServeLogBody() string

	GetServiceClients() []ServerClient

	GetKafkaConsumeConfig() []kafka.KafkaConsumeConfig
	GetKafkaProducerConfig() []kafka.KafkaProductConfig

	GetRedisConfig() []redis.RedisConfig

	GetSQLConfig() []sql.SQLGroupConfig

	GetServiceName() string

	GetMonitorInterval() int

	GetJSONDataLogOption() *JSONDataLogOption

	BusinessLogOff() bool
	RequestBodyLogOff() bool
	Circuit() []inits.CircuitConfig
	Tags() []string
	SuccStatCode() []int
}
