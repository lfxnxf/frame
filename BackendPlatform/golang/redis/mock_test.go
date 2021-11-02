package redis

import (
	"context"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestNewMockRedis(t *testing.T) {
	mockRedis, close, err := NewMockRedis()
	if err != nil {
		return
	}
	defer close()
	_, err = redis.String(mockRedis.DoCtx(context.TODO(), "SET", "foo", "bar"))
	assert.Nil(t, err)

	ret, err := redis.String(mockRedis.DoCtx(context.TODO(), "GET", "foo"))
	assert.Nil(t, err)
	assert.Equal(t, "bar", ret)

	ret, err = redis.String(mockRedis.Do("SETEX", "hello", 1, "world"))
	assert.Equal(t, "OK", ret)
	ttl, err := redis.Int64(mockRedis.Do("TTL", "hello"))
	assert.Equal(t, int64(1), ttl)

}
