package test

import (
	"github.com/TestsLing/aj-captcha-go/service"
	"testing"
)

func TestLocalCacheService_Increment(t *testing.T) {
	cache := service.NewMemCacheService(10)
	key := "test"
	cache.Increment(key, 1)

	if cache.Get(key) != "1" {
		t.Fatal("自增值不正确")
	}

	cache.Increment(key, 2)

	if cache.Get(key) != "3" {
		t.Fatal("自增值不正确")
	}
}
