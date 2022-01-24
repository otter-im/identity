package app

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/otter-im/identity/internal/config"
	"sync"
)

var (
	redisRingOnce  sync.Once
	redisRing      *redis.Ring
	redisCacheOnce sync.Once
	redisCache     *cache.Cache
)

func RedisRing() *redis.Ring {
	redisRingOnce.Do(func() {
		options := &redis.RingOptions{
			Addrs:    config.RedisNodes(),
			Password: config.RedisPassword(),
			DB:       config.RedisDB(),
		}

		redisRing = redis.NewRing(options)
	})
	return redisRing
}

func RedisCache() *cache.Cache {
	redisCacheOnce.Do(func() {
		options := &cache.Options{
			Redis: RedisRing(),
		}
		redisCache = cache.New(options)
	})
	return redisCache
}
