package test

import (
	"github.com/TestsLing/aj-captcha-go/util"
	"testing"
	"time"
)

func TestRedisUtil_Delete(t *testing.T) {
	cache := util.NewRedisUtil()
	val := "testval"
	key := "test"
	cache.Set(key, val, 10)
	cache.Delete(key)

	if cache.Get(key) != "" {
		t.Fatal("缓存删除值失败")
	}
}

func TestRedisUtil_Exists(t *testing.T) {
	cache := util.NewRedisUtil()
	val := "testval"
	key := "test"
	key1 := "test1"
	cache.Set(key, val, 10)
	cache.Set(key1, val, 0)

	if cache.Exists(key) != true {
		t.Fatal("Exists 获取的值不符合要求")
	}

	if cache.Exists(key1) != true {
		t.Fatal("Exists 获取的值不符合要求")
	}
}

func TestRedisUtil_Get(t *testing.T) {
	cache := util.NewRedisUtil()
	val := "testval"
	key := "test"
	cache.Set(key, val, 10)

	if cache.Get(key) != val {
		t.Fatal("获取的值不符合要求")
	}

	time.Sleep(time.Duration(11) * time.Second)

	if cache.Get(key) != "" {
		t.Fatal("时间失效失败")
	}
}
