package recovery

import (
	"testing"

	"github.com/lfxnxf/frame/logic/inits/internal/core"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestRecovery(t *testing.T) {
	assert := assert.New(t)
	c := core.New(nil)
	c.Use(Recovery(true))
	c.Use(core.Function(func(ctx context.Context, c core.Core) {
		panic("test")
	}))
	c.Next(context.TODO())
	assert.NotNil(c.Err())
}
