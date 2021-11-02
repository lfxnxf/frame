package proxy

import (
	"testing"

	"github.com/lfxnxf/frame/BackendPlatform/golang/redis"
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

var (
	mockRedisName = "mock_redis"
	listKey       = "list-key"
	stringKey     = "string-key"
	hashKey       = "hash-key"
	setKey        = "set-key"
	zsetKey       = "zset-key"
)

func TestInitRedis(t *testing.T) {
	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}
}

func TestRedis_RPop(t *testing.T) {
	ns, d := testInitInits("Redis_RPop", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}
	ctx := inits.WithAPPKey(context.Background(), ns)
	num, err := rds.LLen(ctx, "list-key")
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(0), num)

	err = rds.LPush(ctx, listKey, "bar-bar")
	assert.Equal(t, nil, err)

	result, err := rds.RPop(ctx, listKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, "bar-bar", result)

	t.Logf("redis RPOP result:%s", result)
}

func TestRedis_Set(t *testing.T) {
	ns, d := testInitInits("Redis_Set", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}

	ctx := inits.WithAPPKey(context.Background(), ns)
	ok, err := rds.Set(ctx, stringKey, "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	b, err := rds.Get(ctx, stringKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, "bar-bar", string(b[:]))

	byteArr, err := rds.MGet(ctx, stringKey, stringKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, [][]byte{[]byte("bar-bar"), []byte("bar-bar")}, byteArr)

	num, err := rds.Del(ctx, stringKey, "non-key")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, num)

	b, err = rds.Get(ctx, stringKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", string(b[:]))
}

func TestRedis_SetExSecond(t *testing.T) {
	ns, d := testInitInits("Redis_SetExSecond", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}

	ctx := inits.WithAPPKey(context.Background(), ns)
	result, err := rds.SetExSecond(ctx, stringKey, "bar-bar", 2)
	assert.Equal(t, nil, err)
	assert.Equal(t, "OK", result)
	t.Logf("redis SETEX result:%s", result)

	b, err := rds.Get(ctx, stringKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, "bar-bar", string(b[:]))
	t.Logf("redis GET result:%q", b)

	// redis 过期需要手动调用过期逻辑
	// time.Sleep(2 * time.Second)
	// b2, err := rds.Get(context.Background(), stringKey)
	// assert.Equal(t, nil, err)
	// assert.Equal(t, "", string(b2[:]))

}

func TestRedis_HSet(t *testing.T) {
	ns, d := testInitInits("Redis_HSet", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}

	ctx := inits.WithAPPKey(context.Background(), ns)
	num, err := rds.HSet(ctx, hashKey, "bar-bar", "foo")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, num)

	resultMap, err := rds.HGetAll(ctx, hashKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, map[string]string{"bar-bar": "foo"}, resultMap)

	result, err := rds.HGet(ctx, hashKey, "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo", result)

	resultArr, err := rds.HMGet(ctx, hashKey, "bar-bar", "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"foo", "foo"}, resultArr)

	msetResult, err := rds.HMSet(ctx, hashKey, "mon-mon", "loo", "ha-ha", "yoo")
	assert.Equal(t, nil, err)
	assert.Equal(t, "OK", msetResult)

	result, err = rds.HGet(ctx, hashKey, "mon-mon")
	assert.Equal(t, nil, err)
	assert.Equal(t, "loo", result)

	num, err = rds.HDel(ctx, hashKey, "ha-ha", "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, num)

	result, err = rds.HGet(ctx, hashKey, "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, "", result)
}

func TestRedis_SAdd(t *testing.T) {
	ns, d := testInitInits("Redis_SAdd", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}
	ctx := inits.WithAPPKey(context.Background(), ns)
	num, err := rds.SAdd(ctx, setKey, "bar-bar", "foo-foo", "ha-ha", "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, num)

	ok, err := rds.SIsMember(ctx, setKey, "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	ok, err = rds.SIsMember(ctx, setKey, "ha-ha")
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)

	ok, err = rds.SIsMember(ctx, setKey, "ok-ok")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, ok)

	num, err = rds.SRem(ctx, setKey, "foo-foo", "ha-ha")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, num)

	ok, err = rds.SIsMember(ctx, setKey, "ha-ha")
	assert.Equal(t, nil, err)
	assert.Equal(t, false, ok)
}

func TestRedis_ZAdd(t *testing.T) {
	ns, d := testInitInits("Redis_ZAdd", "./testconfigRedis/", "./testconfigRedis/config.toml", "")

	mockRedis, closeFunc, err := redis.NewMockRedis()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	err = d.AddRedisClient(mockRedisName, mockRedis)
	if err != nil {
		panic(err)
	}

	rds := InitRedis(mockRedisName)
	if rds == nil {
		t.Fail()
	} else {
		t.Logf("%+v", *rds)
	}

	ctx := inits.WithAPPKey(context.Background(), ns)
	num, err := rds.ZAdd(ctx, zsetKey, 123, "far-far")
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, num)

	num, err = rds.ZAdd(ctx, zsetKey, 456, "bar-bar", 789, "zar-zar", 101, "kar-kar", 101, "kar-kar", 105, "mar-mar")
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, num)

	num, err = rds.ZCard(ctx, zsetKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, num)

	score, err := rds.ZScore(ctx, zsetKey, "bar-bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, float64(456), score)

	num, err = rds.ZRem(ctx, zsetKey, "bar-bar", "zar-zar")
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, num)

	resultArr, err := rds.ZRange(ctx, zsetKey, 0, 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"kar-kar", "mar-mar"}, resultArr)

	resultArr, err = rds.ZRangeWithScore(ctx, zsetKey, 1, 2)
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"mar-mar", "105", "far-far", "123"}, resultArr)
}
