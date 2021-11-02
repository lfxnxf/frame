package dao

import (
	"context"

	"github.com/lfxnxf/frame/api-gateway/conf"
)

// Dao represents data access object
type Dao struct {
	c *conf.Config
}

func New(c *conf.Config) *Dao {
	return &Dao{
		c: c,
	}
}

// Ping check db resource status
func (d *Dao) Ping(ctx context.Context) error {
	return nil
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}

