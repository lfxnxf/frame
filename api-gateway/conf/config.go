package conf

import (
	"github.com/lfxnxf/frame/logic/inits"
)

type Config struct {
}

func Init() (*Config, error) {
	// parse Config from config file
	cfg := &Config{}
	err := inits.ConfigInstance().Scan(cfg)
	return cfg, err
}
