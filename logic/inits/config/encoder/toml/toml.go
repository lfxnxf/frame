package toml

import (
	"bytes"

	"github.com/lfxnxf/frame/logic/inits/config/encoder"
	"github.com/BurntSushi/toml"
)

type tomlEncoder struct{}

func (t tomlEncoder) Encode(v interface{}) ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	err := toml.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (t tomlEncoder) Decode(d []byte, v interface{}) error {
	return toml.Unmarshal(d, v)
}

func (t tomlEncoder) String() string {
	return "toml"
}

// NewEncoder is a toml encoder
func NewEncoder() encoder.Encoder {
	return tomlEncoder{}
}
