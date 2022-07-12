package util

import (
	"bytes"
	"encoding/base64"
	"github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/model/vo"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

type ImageUtil struct {
	Src       string
	SrcImage  image.Image
	RgbaImage *image.RGBA
	Width     int
	Height    int
}

func NewImageUtil(src string) *ImageUtil {
	srcImage := OpenPngImage(src)

	return &ImageUtil{Src: src,
		SrcImage:  srcImage,
		RgbaImage: ImageToRGBA(srcImage),
		Width:     srcImage.Bounds().Dx(),
		Height:    srcImage.Bounds().Dy(),
	}
}

// IsOpacity 该像素是否透明
func (i *ImageUtil) IsOpacity(x, y int) bool {
	A := i.RgbaImage.RGBAAt(x, y).A

	if float32(A) <= 125 {
		return true
	}
	return false
}

// DecodeImageToFile 将图片转换为新的文件 调试使用
func (i *ImageUtil) DecodeImageToFile() {
	filename := "drawImg.png"
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("创建 %s 失败 %v", filename, err)
	}

	err = png.Encode(file, i.RgbaImage)
	if err != nil {
		log.Printf("png %s Encode 失败 %v", filename, err)
	}

}

// SetText 为图片设置文字
func (i *ImageUtil) SetText(text string, fontsize int, color color.RGBA) {

	x := float64(i.Width) - float64(GetEnOrChLength(text))
	y := float64(i.Height) - (25 / 2) + 7

	font := NewFontUtil(constant.DefaultFont)

	fc := freetype.NewContext()
	// 设置屏幕每英寸的分辨率
	//fc.SetDPI(72)
	// 设置用于绘制文本的字体
	fc.SetFont(font.GetFont())
	// 以磅为单位设置字体大小
	fc.SetFontSize(float64(fontsize))
	// 设置剪裁矩形以进行绘制
	fc.SetClip(i.RgbaImage.Bounds())
	// 设置目标图像
	fc.SetDst(i.RgbaImage)
	// 设置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.NewUniform(color))
	// 设置水印地址
	pt := freetype.Pt(int(x), int(y))
	// 根据 Pt 的坐标值绘制给定的文本内容
	_, err := fc.DrawString(text, pt)
	if err != nil {
		log.Println("构造水印失败:", err)
	}
}

// SetArtText 为图片设置文字
func (i *ImageUtil) SetArtText(text string, fontsize int, point vo.PointVO) {

	font := NewFontUtil(constant.DefaultFont)

	fc := freetype.NewContext()
	// 设置屏幕每英寸的分辨率
	//fc.SetDPI(72)
	// 设置用于绘制文本的字体
	fc.SetFont(font.GetFont())
	// 以磅为单位设置字体大小
	fc.SetFontSize(float64(fontsize))
	// 设置剪裁矩形以进行绘制
	fc.SetClip(i.RgbaImage.Bounds())
	// 设置目标图像
	fc.SetDst(i.RgbaImage)
	// 设置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.NewUniform(color.RGBA{R: uint8(RandomInt(1, 200)), G: uint8(RandomInt(1, 200)), B: uint8(RandomInt(1, 200)), A: 255}))
	// 设置水印地址
	pt := freetype.Pt(point.X, point.Y)
	// 根据 Pt 的坐标值绘制给定的文本内容
	_, err := fc.DrawString(text, pt)
	if err != nil {
		log.Fatalln("构造水印失败:", err)
	}
}

// SetPixel 为像素设置颜色
func (i *ImageUtil) SetPixel(rgba color.RGBA, x, y int) {
	i.RgbaImage.SetRGBA(x, y, rgba)
}

// Base64 为像素设置颜色
func (i *ImageUtil) Base64() string {
	// 开辟一个新的空buff
	var buf bytes.Buffer
	// img写入到buff
	err := png.Encode(&buf, i.RgbaImage)

	if err != nil {
		log.Fatalln("img写入buf失败")
	}

	//开辟存储空间
	dist := make([]byte, buf.Cap()+buf.Len())
	// buff转成base64
	base64.StdEncoding.Encode(dist, buf.Bytes())
	return string(dist)
}

// VagueImage 模糊区域
func (i *ImageUtil) VagueImage(x int, y int) {
	var red uint32
	var green uint32
	var blue uint32
	var alpha uint32

	points := [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, point := range points {
		pointX := x + point[0]
		pointY := y + point[1]

		if pointX < 0 || pointX >= i.Width || pointY < 0 || pointY >= i.Height {
			continue
		}

		r, g, b, a := i.RgbaImage.RGBAAt(pointX, pointY).RGBA()
		red += r >> 8
		green += g >> 8
		blue += b >> 8
		alpha += a >> 8

	}

	var avg uint32
	avg = 8

	rgba := color.RGBA{R: uint8(red / avg), G: uint8(green / avg), B: uint8(blue / avg), A: uint8(alpha / avg)}

	i.RgbaImage.SetRGBA(x, y, rgba)
}

// OpenPngImage 打开png图片
func OpenPngImage(src string) image.Image {
	ff, err := os.Open(src)
	if err != nil {
		log.Printf("打开 %s 图片失败: %v", src, err)
	}

	img, err := png.Decode(ff)

	if err != nil {
		log.Printf("png %s decode  失败: %v", src, err)
	}

	return img
}

// ImageToRGBA 图片转rgba
func ImageToRGBA(img image.Image) *image.RGBA {
	// No conversion needed if image is an *image.RGBA.
	if dst, ok := img.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), img, b.Min, draw.Src)
	return dst
}

func CurrentAbPath() (dir string) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir, _ = filepath.EvalSymlinks(filepath.Dir(exePath))
	tempDir := os.Getenv("TEMP")
	if tempDir == "" {
		tempDir = os.Getenv("TMP")
	}
	tDir, _ := filepath.EvalSymlinks(tempDir)
	if strings.Contains(dir, tDir) {
		//return getCurrentAbPathByCaller()
		var abPath string
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abPath = path.Dir(filename)
		}
		return abPath
	}
	return dir
}
