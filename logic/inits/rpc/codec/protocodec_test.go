package codec

import (
	metadata "github.com/lfxnxf/frame/logic/inits/rpc/internal/metadata"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtoEncode(t *testing.T) {
	codec := NewProtoCodec()
	assert := assert.New(t)

	d := new(metadata.RpcMeta)
	d.Type = metadata.RpcMeta_REQUEST.Enum()
	d.SequenceId = proto.Uint64(1)
	_, err := codec.Encode(d)
	assert.Nil(err)
}

func TestProtoDecode(t *testing.T) {
	codec := NewProtoCodec()
	assert := assert.New(t)

	d := new(metadata.RpcMeta)
	d.Type = metadata.RpcMeta_REQUEST.Enum()
	d.SequenceId = proto.Uint64(1)
	body, err := codec.Encode(d)
	assert.Nil(err)

	p := new(metadata.RpcMeta)
	err = codec.Decode(body, p)
	assert.Nil(err)
}
