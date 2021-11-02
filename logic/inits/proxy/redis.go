package proxy

import (
	"fmt"
	"time"

	"github.com/lfxnxf/frame/logic/inits"
	"golang.org/x/net/context"
)

type Redis struct {
	name string
}

func InitRedis(name string) *Redis {
	return &Redis{name}
}

func (r *Redis) Do(ctx context.Context, cmd string, args ...interface{}) (interface{}, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Do(cmd, args...)
}

func (r *Redis) RPop(ctx context.Context, key string) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.RPop(key)
}

func (r *Redis) LPush(ctx context.Context, name string, fields ...interface{}) error {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return fmt.Errorf("%s redis not exist", r.name)
	}
	return client.LPush(name, fields...)
}

func (r *Redis) LLen(ctx context.Context, key string) (int64, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.LLen(key)
}

func (r *Redis) Send(ctx context.Context, name string, fields ...interface{}) error {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Send(name, fields...)
}

func (r *Redis) ReceiveBlock(ctx context.Context, name string, closech chan struct{}, bufferSize int, block int) chan []byte {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil
	}
	return client.ReceiveBlock(name, closech, bufferSize, block)
}

func (r *Redis) Receive(ctx context.Context, name string, closech chan struct{}, bufferSize int) chan []byte {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil
	}
	return client.Receive(name, closech, bufferSize)
}

func (r *Redis) Set(ctx context.Context, key, value interface{}) (bool, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return false, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Set(key, value)
}

func (r *Redis) SetExSecond(ctx context.Context, key, value interface{}, dur int) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.SetExSecond(key, value, dur)
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Get(key)
}

func (r *Redis) GetInt(ctx context.Context, key string) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.GetInt(key)
}

func (r *Redis) MGet(ctx context.Context, keys ...interface{}) ([][]byte, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.MGet(keys...)
}

func (r *Redis) MSet(ctx context.Context, keys ...interface{}) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.MSet(keys...)
}

func (r *Redis) Del(ctx context.Context, args ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Del(args...)
}

func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return false, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Exists(key)
}

func (r *Redis) Expire(ctx context.Context, key string, expire time.Duration) error {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Expire(key, expire)
}

func (r *Redis) Incrby(ctx context.Context, key string, incr int) (int64, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Incrby(key, incr)
}

func (r *Redis) HDel(ctx context.Context, key interface{}, fields ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HDel(key, fields...)
}

func (r *Redis) HSet(ctx context.Context, key, fieldk string, fieldv interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HSet(key, fieldk, fieldv)
}

func (r *Redis) HGet(ctx context.Context, key, field string) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HGet(key, field)
}

func (r *Redis) HGetInt(ctx context.Context, key, field string) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HGetInt(key, field)
}

func (r *Redis) HMGet(ctx context.Context, key string, fields ...interface{}) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HMGet(key, fields...)
}

func (r *Redis) HMSet(ctx context.Context, key string, fields ...interface{}) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HMSet(key, fields...)
}

func (r *Redis) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HGetAll(key)
}

func (r *Redis) HKeys(ctx context.Context, key string) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HKeys(key)
}

func (r *Redis) HIncrby(ctx context.Context, key, field string, incr int) (int64, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.HIncrby(key, field, incr)
}

func (r *Redis) SAdd(ctx context.Context, key string, members ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.SAdd(key, members...)
}

func (r *Redis) SRem(ctx context.Context, key string, members ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.SRem(key, members...)
}

func (r *Redis) SIsMember(ctx context.Context, key string, member string) (bool, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return false, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.SIsMember(key, member)
}

func (r *Redis) SMembers(ctx context.Context, key string) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.SMembers(key)
}

func (r *Redis) ZAdd(ctx context.Context, key string, args ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZAdd(key, args...)
}

func (r *Redis) ZRange(ctx context.Context, key string, args ...interface{}) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRange(key, args...)
}

func (r *Redis) ZRangeInt(ctx context.Context, key string, start, stop int) ([]int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRangeInt(key, start, stop)
}

func (r *Redis) ZRangeWithScore(ctx context.Context, key string, start, stop int) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRangeWithScore(key, start, stop)
}

func (r *Redis) ZRevRangeWithScore(ctx context.Context, key string, start, stop int) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRevRangeWithScore(key, start, stop)
}

func (r *Redis) ZCount(ctx context.Context, key string, min, max int) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZCount(key, min, max)
}

func (r *Redis) ZCard(ctx context.Context, key string) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZCard(key)
}

func (r *Redis) ZIncrby(ctx context.Context, key string, incr int, member string) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZIncrby(key, incr, member)
}

/*
* If the member not in the zset or key not exits, ZRank will return ErrNil
 */

func (r *Redis) ZRank(ctx context.Context, key string, member string) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRank(key, member)
}

/*
* If the members not in the zset or key not exits, ZRem will return ErrNil
 */

func (r *Redis) ZRem(ctx context.Context, key string, members ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRem(key, members...)
}

func (r *Redis) ZRemrangebyrank(ctx context.Context, key string, members ...interface{}) (int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZRemrangebyrank(key, members...)
}

/*
* If the member not in the zset or key not exits, ZScore will return ErrNil
 */

func (r *Redis) ZScore(ctx context.Context, key, member string) (float64, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return 0, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZScore(key, member)
}

func (r *Redis) Zrevrange(ctx context.Context, key string, args ...interface{}) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Zrevrange(key, args...)
}

func (r *Redis) Zrevrangebyscore(ctx context.Context, key string, args ...interface{}) ([]string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Zrevrangebyscore(key, args...)
}

func (r *Redis) ZrevrangebyscoreInt(ctx context.Context, key string, args ...interface{}) ([]int, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.ZrevrangebyscoreInt(key, args...)
}

func (r *Redis) Subscribe(ctx context.Context, key string, maxSize int) (chan []byte, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return nil, fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Subscribe(ctx, key, maxSize)
}

func (r *Redis) TryLock(ctx context.Context, key string, acquireTimeout, expireTimeout time.Duration) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.TryLock(key, acquireTimeout, expireTimeout)
}

func (r *Redis) Lock(ctx context.Context, key string, expire time.Duration) (string, error) {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return "", fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Lock(key, expire)
}

func (r *Redis) Unlock(ctx context.Context, key string, uuid string) error {
	client := inits.RedisClient(ctx, r.name)
	if client == nil {
		return fmt.Errorf("%s redis not exist", r.name)
	}
	return client.Unlock(key, uuid)
}
