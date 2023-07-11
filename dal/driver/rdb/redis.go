package rdb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"WebService/conf"

	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client
var redisOnce sync.Once

// var rdb
func InitRedis() *redis.Client {
	redisOnce.Do(func() {
		cf := conf.GetConfig()
		redisCli = redis.NewClient(&redis.Options{
			Addr:     cf.Redis.Addr,
			Password: cf.Redis.Passwd, // no password set
			DB:       0,               // use default DB
		})
	})
	return redisCli
}
func GetClient() *redis.Client {
	if redisCli == nil {
		redisCli = InitRedis()
	}
	return redisCli
}

const unlockScript = `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end`

func Set(ctx context.Context, k string, v string, e time.Duration) error {
	cli := GetClient()
	return cli.Set(ctx, k, v, e).Err()
}

func Get(ctx context.Context, k string) (string, error) {
	cli := GetClient()
	val := cli.Get(ctx, k)
	err := val.Err()
	if err == redis.Nil {
		return "", nil
	}
	return val.Result()
}

func Lock(ctx context.Context, k, v string, e time.Duration) error {
	cli := GetClient()
	return cli.SetNX(ctx, k, v, e).Err()
}
func Unlock(ctx context.Context, k, v string) error {
	cli := GetClient()
	cmd := cli.Eval(ctx, unlockScript, []string{k}, v)
	if e := cmd.Err(); e != nil {
		return e
	}
	if cmd.Val() == 0 {
		return fmt.Errorf("Unlock err:key-%s val-%s", k, v)
	}
	return nil
}
