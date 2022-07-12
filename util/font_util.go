package util

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
	"log"
	"path/filepath"
	"unicode"
)

type FontUtil struct {
	Src string
}

func NewFontUtil(src string) *FontUtil {
	return &FontUtil{Src: src}
}

// GetFont 获取一个字体对象
func (f *FontUtil) GetFont() *truetype.Font {
	root := filepath.Dir(CurrentAbPath())
	fontSourceBytes, err := ioutil.ReadFile(root + f.Src)
	if err != nil {
		log.Println("读取字体失败:", err)
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)

	if err != nil {
		log.Println("解析字体失败:", err)
	}

	return trueTypeFont
}

func GetEnOrChLength(text string) int {
	enCount, zhCount := 0, 0

	for _, t := range text {
		if unicode.Is(unicode.Han, t) {
			zhCount++
		} else {
			enCount++
		}
	}

	chOffset := (25/2)*zhCount + 5
	enOffset := enCount * 8

	return chOffset + enOffset
}
