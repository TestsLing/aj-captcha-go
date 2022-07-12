package service

import (
	"github.com/TestsLing/aj-captcha-go/util"
	"strconv"
)

type MemCacheService struct {
	Cache *util.CacheUtil
}

func NewMemCacheService(captchaCacheMaxNumber int) CaptchaCacheInterface {
	return &MemCacheService{Cache: util.NewCacheUtil(captchaCacheMaxNumber)}
}

func (l *MemCacheService) Get(key string) string {
	return l.Cache.Get(key)
}

func (l *MemCacheService) Set(key string, val string, expiresInSeconds int) {
	l.Cache.Set(key, val, expiresInSeconds)
}

func (l *MemCacheService) Delete(key string) {
	l.Cache.Delete(key)
}

func (l *MemCacheService) Exists(key string) bool {
	return l.Cache.Exists(key)
}

func (l *MemCacheService) GetType() string {
	return "mem"
}

func (l *MemCacheService) Increment(key string, val int) int {
	cacheVal := l.Cache.Get(key)
	num, err := strconv.Atoi(cacheVal)
	if err != nil {
		num = 0
	}

	ret := num + val

	l.Cache.Set(key, strconv.Itoa(ret), 0)

	return ret
}
