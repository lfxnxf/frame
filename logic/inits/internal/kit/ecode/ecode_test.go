package ecode

import (
	"context"
	"testing"

	"github.com/lfxnxf/frame/logic/inits/ratelimit"
	"github.com/stretchr/testify/assert"
)

func TestConvertHttpStatus(t *testing.T) {
	assert.Equal(t, ConvertHttpStatus(ratelimit.ErrLimited), 501)
	assert.Equal(t, ConvertHttpStatus(nil), 200)
	assert.Equal(t, ConvertHttpStatus(context.DeadlineExceeded), 500)
}
