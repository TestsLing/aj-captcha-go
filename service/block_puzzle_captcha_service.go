package service

import (
	"encoding/json"
	"errors"
	"fmt"
	constant "github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/model/vo"
	"github.com/TestsLing/aj-captcha-go/util"
	img "github.com/TestsLing/aj-captcha-go/util/image"
	"golang.org/x/image/colornames"
	"log"
	"math"
)

type BlockPuzzleCaptchaService struct {
	point   vo.PointVO
	factory *CaptchaServiceFactory
}

func NewBlockPuzzleCaptchaService(factory *CaptchaServiceFactory) *BlockPuzzleCaptchaService {
	return &BlockPuzzleCaptchaService{
		factory: factory,
	}
}

// Get 获取验证码图片信息
func (b *BlockPuzzleCaptchaService) Get() map[string]interface{} {

	// 初始化背景图片
	backgroundImage := img.GetBackgroundImage()

	// 为背景图片设置水印
	backgroundImage.SetText(b.factory.config.Watermark.Text, b.factory.config.Watermark.FontSize, b.factory.config.Watermark.Color)

	// 初始化模板图片
	templateImage := img.GetTemplateImage()

	// 构造前端所需图片
	b.pictureTemplatesCut(backgroundImage, templateImage)

	data := make(map[string]interface{})
	data["originalImageBase64"] = backgroundImage.Base64()
	data["jigsawImageBase64"] = templateImage.Base64()
	data["secretKey"] = b.point.SecretKey
	data["token"] = util.GetUuid()

	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, data["token"])
	jsonPoint, err := json.Marshal(b.point)
	if err != nil {
		log.Fatalln("point json err:", err)
	}

	b.factory.GetCache().Set(codeKey, string(jsonPoint), b.factory.config.CacheExpireSec)

	return data
}

func (b *BlockPuzzleCaptchaService) pictureTemplatesCut(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil) {
	// 生成拼图坐标点
	b.generateJigsawPoint(backgroundImage, templateImage)
	// 裁剪模板图
	b.cutByTemplate(backgroundImage, templateImage, b.point.X, 0)

	// 插入干扰图
	for {
		newTemplateImage := img.GetTemplateImage()
		if newTemplateImage.Src != templateImage.Src {
			offsetX := util.RandomInt(0, backgroundImage.Width-newTemplateImage.Width-5)
			if math.Abs(float64(newTemplateImage.Width-offsetX)) > float64(newTemplateImage.Width/2) {
				b.interferenceByTemplate(backgroundImage, newTemplateImage, offsetX, b.point.Y)
				break
			}
		}
	}
}

// 插入干扰图
func (b *BlockPuzzleCaptchaService) interferenceByTemplate(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil, x1 int, y1 int) {
	xLength := templateImage.Width
	yLength := templateImage.Height

	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			// 如果模板图像当前像素点不是透明色 copy源文件信息到目标图片中
			isOpacity := templateImage.IsOpacity(x, y)

			// 当前模板像素在背景图中的位置
			backgroundX := x + x1
			backgroundY := y + y1

			// 当不为透明时
			if !isOpacity {
				// 背景图区域模糊
				backgroundImage.VagueImage(backgroundX, backgroundY)
			}

			//防止数组越界判断
			if x == (xLength-1) || y == (yLength-1) {
				continue
			}

			rightOpacity := templateImage.IsOpacity(x+1, y)
			downOpacity := templateImage.IsOpacity(x, y+1)

			//描边处理，,取带像素和无像素的界点，判断该点是不是临界轮廓点,如果是设置该坐标像素是白色
			if (isOpacity && !rightOpacity) || (!isOpacity && rightOpacity) || (isOpacity && !downOpacity) || (!isOpacity && downOpacity) {
				backgroundImage.RgbaImage.SetRGBA(backgroundX, backgroundY, colornames.White)
			}
		}
	}
}

func (b *BlockPuzzleCaptchaService) cutByTemplate(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil, x1, y1 int) {
	xLength := templateImage.Width
	yLength := templateImage.Height

	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			// 如果模板图像当前像素点不是透明色 copy源文件信息到目标图片中
			isOpacity := templateImage.IsOpacity(x, y)

			// 当前模板像素在背景图中的位置
			backgroundX := x + x1
			backgroundY := y + y1

			// 当不为透明时
			if !isOpacity {
				// 获取原图像素
				backgroundRgba := backgroundImage.RgbaImage.RGBAAt(backgroundX, backgroundY)
				// 将原图的像素扣到模板图上
				templateImage.SetPixel(backgroundRgba, x, y)
				// 背景图区域模糊
				backgroundImage.VagueImage(backgroundX, backgroundY)
			}

			//防止数组越界判断
			if x == (xLength-1) || y == (yLength-1) {
				continue
			}

			rightOpacity := templateImage.IsOpacity(x+1, y)
			downOpacity := templateImage.IsOpacity(x, y+1)

			//描边处理，,取带像素和无像素的界点，判断该点是不是临界轮廓点,如果是设置该坐标像素是白色
			if (isOpacity && !rightOpacity) || (!isOpacity && rightOpacity) || (isOpacity && !downOpacity) || (!isOpacity && downOpacity) {
				templateImage.RgbaImage.SetRGBA(x, y, colornames.White)
				backgroundImage.RgbaImage.SetRGBA(backgroundX, backgroundY, colornames.White)
			}
		}
	}
}

// 生成模板图在背景图中的随机坐标点
func (b *BlockPuzzleCaptchaService) generateJigsawPoint(backgroundImage *util.ImageUtil, templateImage *util.ImageUtil) {
	widthDifference := backgroundImage.Width - templateImage.Width
	heightDifference := backgroundImage.Height - templateImage.Height

	x, y := 0, 0

	if widthDifference <= 0 {
		x = 5
	} else {
		x = util.RandomInt(100, widthDifference-100)
	}
	if heightDifference <= 0 {
		y = 5
	} else {
		y = util.RandomInt(5, heightDifference)
	}
	point := vo.PointVO{X: x, Y: y}
	point.SetSecretKey(util.RandString(16))
	b.point = point
}

func (b *BlockPuzzleCaptchaService) Check(token string, pointJson string) error {
	cache := b.factory.GetCache()

	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, token)

	cachePointInfo := cache.Get(codeKey)

	if cachePointInfo == "" {
		return errors.New("验证码已失效")
	}

	// 解析结构体
	cachePoint := &vo.PointVO{}
	userPoint := &vo.PointVO{}
	err := json.Unmarshal([]byte(cachePointInfo), cachePoint)

	if err != nil {
		return err
	}

	// 解密前端传递过来的数据
	userPointJson := util.AesDecrypt(pointJson, cachePoint.SecretKey)

	err = json.Unmarshal([]byte(userPointJson), userPoint)

	if err != nil {
		return err
	}

	// 校验两个点是否符合
	if math.Abs(float64(cachePoint.X-userPoint.X)) <= float64(b.factory.config.BlockPuzzle.Offset) && cachePoint.Y == userPoint.Y {
		return nil
	}

	return errors.New("验证失败")
}

func (b *BlockPuzzleCaptchaService) Verification(token string, pointJson string) error {
	err := b.Check(token, pointJson)
	if err != nil {
		return err
	}
	codeKey := fmt.Sprintf(constant.CodeKeyPrefix, token)
	b.factory.GetCache().Delete(codeKey)
	return nil
}
