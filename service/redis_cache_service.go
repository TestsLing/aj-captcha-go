package service

import (
	"github.com/TestsLing/aj-captcha-go/util"
	"strconv"
)

type RedisCacheService struct {
	Cache *util.RedisUtil
}

func NewRedisCacheService() CaptchaCacheInterface {
	redisUtils := util.NewRedisUtil()
	return &RedisCacheService{Cache: redisUtils}
}

func (l *RedisCacheService) Get(key string) string {
	return l.Cache.Get(key)
}

func (l *RedisCacheService) Set(key string, val string, expiresInSeconds int) {
	l.Cache.Set(key, val, expiresInSeconds)
}

func (l *RedisCacheService) Delete(key string) {
	l.Cache.Delete(key)
}

func (l *RedisCacheService) Exists(key string) bool {
	return l.Cache.Exists(key)
}

func (l *RedisCacheService) GetType() string {
	return "redis"
}

func (l *RedisCacheService) Increment(key string, val int) int {
	cacheVal := l.Cache.Get(key)
	num, err := strconv.Atoi(cacheVal)
	if err != nil {
		num = 0
	}

	ret := num + val

	l.Cache.Set(key, strconv.Itoa(ret), 0)
	return ret
}
