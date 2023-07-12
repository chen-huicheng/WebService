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

// InitRedis 初始化 Redis
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

// GetClient 使用 redis 全局变量
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

// Lock 加锁
func Lock(ctx context.Context, key, token string, e time.Duration) (bool, error) {
	cli := GetClient()
	ok, err := cli.SetNX(ctx, key, token, e).Result()
	if err != nil {
		return cli.SetNX(ctx, key, token, e).Result()
	}
	return ok, nil
}

// Unlock 解锁
/*
	验证身份并解锁，需要执行两次Redis操作，使用 lua script 脚本保证原子性
*/
func Unlock(ctx context.Context, key, token string) error {
	cli := GetClient()
	cmd := cli.Eval(ctx, unlockScript, []string{key}, token)
	if e := cmd.Err(); e != nil {
		return e
	}
	if cmd.Val() == 0 {
		return fmt.Errorf("Unlock err:key-%s val-%s", key, token)
	}
	return nil
}

// MaxLockExpire 最长加锁超时时间 仅限续约生效
const MaxLockExpire = time.Minute * 1

// AutoRefresh 续约机制
func AutoRefresh(ctx context.Context, key, token string, e time.Duration, interval time.Duration) {
	go func() {
		cli := GetClient()
		ticker := time.NewTicker(interval)
		endTicker := time.NewTicker(MaxLockExpire)
		for {
			select {
			case <-ticker.C:
				val, err := cli.Get(ctx, key).Result()
				if err != nil || val != token { //判断是否是该用户加的锁
					return
				}
				_, err = cli.Expire(ctx, key, e).Result() //修改过期时间
				if err != nil {
					return
				}
			case <-endTicker.C: // 锁持有的最长时间，超过则不再续约
				return
			}
		}
	}()
}

// LockAndAutoRefresh 加锁并自动续约
/*
	key:锁  token:锁持有者才知道 e:超时时间 必须设置避免锁持有协程异常退出造成死锁
	interval:续约间隔
*/
func LockAndAutoRefresh(ctx context.Context, key, token string, e time.Duration, interval time.Duration) (bool, error) {
	cli := GetClient()
	ok, err := cli.SetNX(ctx, key, token, e).Result()
	if err != nil {
		ok, err = cli.SetNX(ctx, key, token, e).Result()
	}
	if ok {
		AutoRefresh(ctx, key, token, e, interval)
	}
	return ok, err
}
