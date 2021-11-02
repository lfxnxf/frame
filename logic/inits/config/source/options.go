package source

import (
	"github.com/lfxnxf/frame/logic/inits/config/encoder"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/toml"
	"golang.org/x/net/context"
)

type Options struct {
	Encoder encoder.Encoder
	Context context.Context
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: toml.NewEncoder(),
		Context: context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}
