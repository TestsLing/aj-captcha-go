package test

import (
	"fmt"
	"github.com/TestsLing/aj-captcha-go/service"
	"github.com/TestsLing/aj-captcha-go/util"
	"image/color"
	"path/filepath"
	"testing"
)

func TestBlockPuzzleCaptchaService_Get(t *testing.T) {
	//
	//vo := &vo2.CaptchaVO{}
	//b := &service.BlockPuzzleCaptchaService{}
	//res := b.Get(*vo)
	//
	//fmt.Println(res)
}

func TestImage(t *testing.T) {

	backgroundImage := util.NewImageUtil(filepath.Dir(util.CurrentAbPath()) + "/resources/defaultImages/jigsaw/original/1.png")
	// 为背景图片设置水印
	backgroundImage.SetText("牛逼AAA", 14, color.RGBA{R: 120, G: 120, B: 255, A: 255})
	backgroundImage.DecodeImageToFile()
}

func TestIntCovert(t *testing.T) {

	cache := service.NewMemCacheService(10)

	cache.Set("test1", "tes111", 0)

	val := cache.Get("test1")

	fmt.Println(val)

}
