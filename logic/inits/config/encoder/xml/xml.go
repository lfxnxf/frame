package xml

import (
	"encoding/xml"

	"github.com/lfxnxf/frame/logic/inits/config/encoder"
)

type xmlEncoder struct{}

func (x xmlEncoder) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (x xmlEncoder) Decode(d []byte, v interface{}) error {
	return xml.Unmarshal(d, v)
}

func (x xmlEncoder) String() string {
	return "xml"
}

// NewEncoder is xml encoder
func NewEncoder() encoder.Encoder {
	return xmlEncoder{}
}
