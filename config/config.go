package config

import (
	"github.com/TestsLing/aj-captcha-go/const"
	"image/color"
)

// WatermarkConfig 水印设置
type WatermarkConfig struct {
	FontSize int        `yaml:"fontSize"`
	Color    color.RGBA `yaml:"color"`
	Text     string     `yaml:"text"`
}

type BlockPuzzleConfig struct {
	// 校验时 容错偏移量
	Offset int `yaml:"offset"`
}

type ClickWordConfig struct {
	FontSize int `yaml:"fontSize"`
	FontNum  int `yaml:"fontNum"`
}

// RedisConfig redis配置选项
type RedisConfig struct {
	//redis单机或者集群访问地址
	DBAddress []string `yaml:"dbAddress"`
	//最大空闲连接数
	DBMaxIdle int `yaml:"dbMaxIdle"`
	//最大连接数
	DBMaxActive int `yaml:"dbMaxActive"`
	//redis表示空闲连接保活时间
	DBIdleTimeout int `yaml:"DBIdleTimeout"`
	//redis密码
	DBPassWord string `yaml:"dbPassWord"`
	//是否使用redis集群
	EnableCluster bool `yaml:"enableCluster"`
	//单机模式下使用redis的指定库，比如：0，1，2，3等等，默认为0
	DB int `yaml:"db"`
}

type Config struct {
	Watermark   *WatermarkConfig   `yaml:"watermark"`
	ClickWord   *ClickWordConfig   `yaml:"clickWord"`
	BlockPuzzle *BlockPuzzleConfig `yaml:"blockPuzzle"`
	// 验证码使用的缓存类型
	CacheType      string `yaml:"cacheType"`
	CacheExpireSec int    `yaml:"cacheExpireSec"`
	//redis配置
	Redis *RedisConfig `yaml:"redis"`
	// 项目的绝对路径: 图片、字体等
	ResourcePath string `yaml:"resourcePath"`
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
		ResourcePath: "./",
	}
}
