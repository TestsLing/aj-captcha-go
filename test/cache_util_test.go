package test

import (
	"github.com/TestsLing/aj-captcha-go/util"
	"testing"
	"time"
)

func TestCacheUtil_Exists(t *testing.T) {
	cache := util.NewCacheUtil(10)
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

func TestCacheUtil_Get(t *testing.T) {
	cache := util.NewCacheUtil(10)
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

func TestCacheUtil_Delete(t *testing.T) {
	cache := util.NewCacheUtil(10)
	val := "testval"
	key := "test"
	cache.Set(key, val, 10)
	cache.Delete(key)

	if cache.Get(key) != "" {
		t.Fatal("缓存删除值失败")
	}

}

func TestCacheUtil_MaxSize(t *testing.T) {
	cache := util.NewCacheUtil(2)
	val := "testval"
	key := "test"
	key1 := "test1"
	key2 := "test2"
	cache.Set(key, val, 10)
	cache.Set(key1, val, 10)

	if cache.Get(key) != val {
		t.Fatal("缓存数据异常")
	}

	cache.Set(key2, val, 10)

	if cache.Get(key) != "" {
		t.Fatal("清空缓存数据失败")
	}
	if cache.Get(key2) != val {
		t.Fatal("清空缓存后写入数据失败")
	}

}

func TestCacheUtil_Clear(t *testing.T) {
	cache := util.NewCacheUtil(10)
	val := "testval"
	key := "test"
	key1 := "test1"
	cache.Set(key, val, 10)
	cache.Set(key1, val, 10)
	cache.Clear()

	if len(cache.Data) != 0 {
		t.Fatal("清空缓存失败")
	}
}
