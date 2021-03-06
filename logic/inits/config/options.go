package config

import (
	"github.com/lfxnxf/frame/logic/inits/config/encoder"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/json"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/toml"
	"github.com/lfxnxf/frame/logic/inits/config/source"
)

// Option represents a func
type Option func(o *Options)

// Options represents a option on the source
type Options struct {
	Source  []source.Source
}

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}

// TomlEncoder represents a toml encoder
func TomlEncoder() encoder.Encoder {
	return toml.NewEncoder()
}

// JSONEncoder represents a toml encoder
func JSONEncoder() encoder.Encoder {
	return json.NewEncoder()
}