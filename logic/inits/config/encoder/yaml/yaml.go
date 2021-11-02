package yaml

import (
	"github.com/lfxnxf/frame/logic/inits/config/encoder"
	"github.com/ghodss/yaml"
)

type yamlEncoder struct{}

func (y yamlEncoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (y yamlEncoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

func (y yamlEncoder) String() string {
	return "yaml"
}

// NewEncoder is yaml encoder
func NewEncoder() encoder.Encoder {
	return yamlEncoder{}
}
