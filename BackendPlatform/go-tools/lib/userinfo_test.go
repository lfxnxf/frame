package lib

import (
	"testing"
	"github.com/lfxnxf/frame/logic/rpc-go"
	"golang.org/x/net/context"
	"github.com/stretchr/testify/assert"
	"time"
)

type Config struct {
}

func TestGetUserInfo(t *testing.T) {
	path := "./conf/test.toml"
	var c Config
	_, err := rpc.NewConfigToml(path, &c)
	if err != nil {
		return
	}
	uids := []int64{218313629, 564238347}
	start := 200000
	count := 1000
	for i := start; i < start+count; i++ {
		uids = append(uids, int64(i))
	}
	now := time.Now()
	userinfos, err := GetUserInfo(context.TODO(), uids)
	assert.Equal(t, len(userinfos), count+2)
	assert.Nil(t, err)
	t.Log("cost time:", time.Since(now))
}
