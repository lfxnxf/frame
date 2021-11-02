package manager

import (
	"context"

	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/conf"
)

// Manager represents middleware component
// such as, kafka, http client or rpc client, etc.
type Manager struct {
	c *conf.Config
}

func New(conf *conf.Config) *Manager {
	return &Manager{
		c: conf,
	}
}


// Ping check middleware resource status
func (m *Manager) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (m *Manager) Close() error {
	return nil
}

