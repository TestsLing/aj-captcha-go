package service

import (
	configIns "github.com/TestsLing/aj-captcha-go/config"
	"log"
	"sync"
)

// CaptchaServiceFactory 验证码服务工厂
type CaptchaServiceFactory struct {
	config      *configIns.Config
	ServiceMap  map[string]CaptchaInterface
	ServiceLock sync.RWMutex

	CacheMap  map[string]CaptchaCacheInterface
	CacheLock sync.RWMutex
}

func NewCaptchaServiceFactory(config *configIns.Config) *CaptchaServiceFactory {

	factory := &CaptchaServiceFactory{
		ServiceMap: make(map[string]CaptchaInterface),
		CacheMap:   make(map[string]CaptchaCacheInterface),
		config:     config,
	}
	return factory
}

func (c *CaptchaServiceFactory) GetCache() CaptchaCacheInterface {
	key := c.config.CacheType
	c.CacheLock.RLock()
	defer c.CacheLock.RUnlock()
	if _, ok := c.CacheMap[key]; !ok {
		log.Fatalf("未注册%s类型的Cache", key)
	}
	return c.CacheMap[key]
}

func (c *CaptchaServiceFactory) RegisterCache(key string, cacheInterface CaptchaCacheInterface) {
	c.CacheLock.Lock()
	defer c.CacheLock.Unlock()
	c.CacheMap[key] = cacheInterface
}

func (c *CaptchaServiceFactory) RegisterService(key string, service CaptchaInterface) {
	c.ServiceLock.Lock()
	defer c.ServiceLock.Unlock()
	c.ServiceMap[key] = service
}

func (c *CaptchaServiceFactory) GetService(key string) CaptchaInterface {
	c.ServiceLock.RLock()
	defer c.ServiceLock.RUnlock()
	if _, ok := c.ServiceMap[key]; !ok {
		log.Fatalf("未注册%s类型的Service", key)
	}
	return c.ServiceMap[key]
}
