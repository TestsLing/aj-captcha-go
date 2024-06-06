package image

import (
	"github.com/xierui921326/aj-captcha-go/const"
	"github.com/xierui921326/aj-captcha-go/util"
	"log"
	"os"
	"path/filepath"
)

var backgroundImageArr []string
var clickBackgroundImageArr []string
var templateImageArr []string

var resourceAbsPath string

func SetUp(resourcePath string) {
	resourceAbsPath = resourcePath
	root := resourcePath

	//root := "/Users/skyline/go/src/aj-captcha-go"
	backgroundImageRoot := root + constant.DefaultBackgroundImageDirectory
	templateImageRoot := root + constant.DefaultTemplateImageDirectory
	clickBackgroundImageRoot := root + constant.DefaultClickBackgroundImageDirectory

	err := filepath.Walk(backgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		backgroundImageArr = append(backgroundImageArr, path)
		return nil
	})

	err = filepath.Walk(templateImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		templateImageArr = append(templateImageArr, path)
		return nil
	})

	err = filepath.Walk(clickBackgroundImageRoot, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		clickBackgroundImageArr = append(clickBackgroundImageArr, path)
		return nil
	})

	if err != nil {
		log.Printf("初始化resource目录失败，请检查该目录是否存在 err: %v", err)
	}

}

func GetBackgroundImage() *util.ImageUtil {
	max := len(backgroundImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(backgroundImageArr[util.RandomInt(0, max)], resourceAbsPath+constant.DefaultFont)
}

func GetTemplateImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(templateImageArr[util.RandomInt(0, max)], resourceAbsPath+constant.DefaultFont)
}

func GetClickBackgroundImage() *util.ImageUtil {
	max := len(templateImageArr) - 1
	if max <= 0 {
		max = 1
	}
	return util.NewImageUtil(clickBackgroundImageArr[util.RandomInt(0, max)], resourceAbsPath+constant.DefaultFont)
}
