# Aj-Captcha-Go

## Introduction

`aj-captcha` 滑动验证码的 `Go` 语言实现

### Summary

| 术语    | 描述                                                                  |
|-------|---------------------------------------------------------------------|
| 验证码类型 | 1）滑动拼图 blockPuzzle  2）文字点选 clickWord                                |
| 验证    | 用户拖动/点击一次验证码拼图即视为一次“验证”，不论拼图/点击是否正确                                 |
| 二次校验  | 验证数据随表单提交到后台后，后台需要调用captchaService.verification做二次校验。目的是核实验证数据的有效性。 |

### Features

- 支持滑动拼团验证
- 支持文字点选验证

### 预览效果

![block](https://gitee.com/anji-plus/captcha/raw/master/images/%E6%BB%91%E5%8A%A8%E6%8B%BC%E5%9B%BE.gif)

![click](https://gitee.com/anji-plus/captcha/raw/master/images/%E7%82%B9%E9%80%89%E6%96%87%E5%AD%97.gif)

## Requirements

- github.com/golang/freetype  作为字体绘制依赖
- golang.org/x/image  图片操作

## Configuration

```go
// WatermarkConfig 水印设置
type WatermarkConfig struct {
	FontSize int   // 水印字体大小
	Color    color.RGBA  // 水印rgba颜色
	Text     string // 水印文字
}

type BlockPuzzleConfig struct {
	Offset int // 校验时 容错偏移量
}

type ClickWordConfig struct {
	FontSize int // 点击验证文字的大小
	FontNum  int // 点击验证的文字的随机数量 
}

type Config struct {
	Watermark      *WatermarkConfig
	ClickWord      *ClickWordConfig
	BlockPuzzle    *BlockPuzzleConfig
	CacheType      string // 验证码使用的缓存类型
	CacheExpireSec int
}

func NewConfig() *Config {
	return &Config{
		CacheType: "redis",  // 注册的缓存类型
		Watermark: &WatermarkConfig{
			FontSize: 12,
			Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
			Text:     "我的水印",
		},
		ClickWord: &ClickWordConfig{
			FontSize: 25,
			FontNum:  5,
		},
		BlockPuzzle:    &BlockPuzzleConfig{Offset: 10},
		CacheExpireSec: 2 * 60, // 缓存有效时间
	}
}
```

## Installation

```bash
go get -u github.com/TestsLing/aj-captcha-go
```

## Usage

- gin框架

```go
package main

import (
	config2 "github.com/TestsLing/aj-captcha-go/config"
	"github.com/TestsLing/aj-captcha-go/service"
	"github.com/TestsLing/aj-captcha-go/const"
	"github.com/gin-gonic/gin"
)

// 客户端参数 看自身业务构建即可
type clientParams struct {
	Token       string `json:"token"`
	PointJson   string `json:"pointJson"`
	CaptchaType string `json:"captchaType"`
}

// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
var config = config2.NewConfig()
// 服务工厂，主要用户注册 获取 缓存和验证服务
var factory = service.NewCaptchaServiceFactory(config)

func main() {

	// 这里默认是注册了 内存缓存，但是不足以应对生产环境，希望自行注册缓存驱动 实现缓存接口即可替换（CacheType就是注册进去的 key）
	factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(20)) // 这里20指的是缓存阈值
	
	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))

	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/captcha/get", func(c *gin.Context) {
		// 根据参数类型获取不同服务即可
		data := factory.GetService(constant.BlockPuzzleCaptcha).Get()
		//输出json结果给调用方
		c.JSON(200, data)
	})
	r.Run("0.0.0.0:888") // listen and serve on 0.0.0.0:888
}

```

- http/net

直接运行 `example` 即可

## Changelog

- 2022.7.12  初次提交Go实现

## FAQ

## Support

### Contact

- 微信: hack_mess

## Authors and acknowledgment

## License

- MIT