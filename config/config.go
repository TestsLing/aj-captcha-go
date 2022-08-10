package config

import (
	"github.com/TestsLing/aj-captcha-go/const"
	"image/color"
)

// WatermarkConfig 水印设置
type WatermarkConfig struct {
	FontSize int
	Color    color.RGBA
	Text     string
}

type BlockPuzzleConfig struct {
	Offset int // 校验时 容错偏移量
}

type ClickWordConfig struct {
	FontSize int
	FontNum  int
}

// RedisConfig redis配置选项
type RedisConfig struct {
	//redis单机或者集群访问地址
	DBAddress []string
	//最大空闲连接数
	DBMaxIdle int
	//最大连接数
	DBMaxActive int
	//redis表示空闲连接保活时间
	DBIdleTimeout int
	//redis密码
	DBPassWord string
	//是否使用redis集群
	EnableCluster bool
	//单机模式下使用redis的指定库，比如：0，1，2，3等等，默认为0
	DB int
}

type Config struct {
	Watermark      *WatermarkConfig
	ClickWord      *ClickWordConfig
	BlockPuzzle    *BlockPuzzleConfig
	CacheType      string // 验证码使用的缓存类型
	CacheExpireSec int
	Redis          *RedisConfig //redis配置
}

func NewConfig() *Config {
	return &Config{
		//可以为redis类型缓存RedisCacheKey，也可以为内存MemCacheKey
		CacheType: constant.MemCacheKey,
		Watermark: &WatermarkConfig{
			FontSize: 12,
			Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			Text:     "我的水印",
		},
		ClickWord: &ClickWordConfig{
			FontSize: 25,
			FontNum:  4,
		},
		BlockPuzzle:    &BlockPuzzleConfig{Offset: 10},
		CacheExpireSec: 2 * 60, // 缓存有效时间
		//redis配置选项
		Redis: &RedisConfig{
			DBAddress:     []string{"127.0.0.1:6379"},
			DBPassWord:    "",
			EnableCluster: false,
			DB:            0,
		},
	}
}
