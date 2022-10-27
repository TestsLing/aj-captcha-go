package test

import (
	"github.com/TestsLing/aj-captcha-go/util"
	"testing"
	"time"
)

func TestDftRedisUtil_Delete(t *testing.T) {
	cache := util.NewDftRedisUtil()
	val := "testval"
	key := "testdft"
	cache.Set(key, val, 10)
	cache.Delete(key)

	if cache.Get(key) != "" {
		t.Fatal("default 缓存删除值失败")
	}
}

func TestConfigRedisUtil_Delete(t *testing.T) {
	cache := util.NewConfigRedisUtil([]string{"192.168.1.111:6379"}, "", false, 0)
	val := "testvalconfig"
	key := "testconfig"
	cache.Set(key, val, 10)
	cache.Delete(key)

	if cache.Get(key) != "" {
		t.Fatal("config 缓存删除值失败")
	}
}

func TestDftRedisUtil_Exists(t *testing.T) {
	cache := util.NewDftRedisUtil()
	val := "testvaldft"
	key := "testdft"
	key1 := "test1dft"
	cache.Set(key, val, 10)
	cache.Set(key1, val, 0)

	if cache.Exists(key) != true {
		t.Fatal("default Exists 获取的值不符合要求")
	}

	if cache.Exists(key1) != true {
		t.Fatal("default Exists 获取的值不符合要求")
	}
}

func TestConfigRedisUtil_Exists(t *testing.T) {
	cache := util.NewConfigRedisUtil([]string{"192.168.1.111:6379"}, "", false, 0)
	val := "testvalconfig"
	key := "testconfig"
	key1 := "test1config"
	cache.Set(key, val, 10)
	cache.Set(key1, val, 0)

	if cache.Exists(key) != true {
		t.Fatal("config Exists 获取的值不符合要求")
	}

	if cache.Exists(key1) != true {
		t.Fatal("config Exists 获取的值不符合要求")
	}
}

func TestDftRedisUtil_Get(t *testing.T) {
	cache := util.NewDftRedisUtil()
	val := "testvaldft"
	key := "testdft"
	cache.Set(key, val, 10)

	if cache.Get(key) != val {
		t.Fatal("default 获取的值不符合要求")
	}

	time.Sleep(time.Duration(11) * time.Second)

	if cache.Get(key) != "" {
		t.Fatal("default 时间失效失败")
	}
}

func TestConfigRedisUtil_Get(t *testing.T) {
	cache := util.NewConfigRedisUtil([]string{"192.168.1.111:6379"}, "", false, 0)
	val := "testvalconfig"
	key := "testconfig"
	cache.Set(key, val, 10)

	if cache.Get(key) != val {
		t.Fatal("config 获取的值不符合要求")
	}

	time.Sleep(time.Duration(11) * time.Second)

	if cache.Get(key) != "" {
		t.Fatal("config 时间失效失败")
	}
}
