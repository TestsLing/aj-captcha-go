# Aj-Captcha-Go

## Introduction

`aj-captcha` 滑动验证码的 `Go` 语言实现

### Summary - 概要

| 术语    | 描述                                                                  |
|-------|---------------------------------------------------------------------|
| 验证码类型 | 1）滑动拼图 blockPuzzle  2）文字点选 clickWord                                |
| 验证    | 用户拖动/点击一次验证码拼图即视为一次“验证”，不论拼图/点击是否正确                                 |
| 二次校验  | 验证数据随表单提交到后台后，后台需要调用captchaService.verification做二次校验。目的是核实验证数据的有效性。 |

### Features - 特性

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

## Installation - 安装（如何安装。）

```bash
go download github.com/TestsLing/aj-captcha-go
```

## Usage - 用法（用法。）

```go

```

## Changelog

- 2022.7.12  初次提交Go实现

## FAQ

## Support

### Contact

- 微信: hack_mess

## Authors and acknowledgment

## License

- MIT