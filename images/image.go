package images

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"golang.org/x/image/font"

	"github.com/golang/freetype"
	"github.com/nfnt/resize"
)

//Image 图片自定义处理类
type Image struct {
	m *image.RGBA
}

//NewImage 创建指定大小的画板
func NewImage(w int, h int) (img *Image) {
	img = &Image{}
	img.m = image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img.m, img.m.Bounds(), image.White, image.ZP, draw.Src)
	return
}

//NewImageFromFile 根据文件创建画板
func NewImageFromFile(p string) (img *Image, err error) {
	ig, err := decodeImage(p)
	if err != nil {
		return
	}
	img = &Image{}
	img.m = image.NewRGBA(ig.Bounds())
	draw.Draw(img.m, img.m.Bounds(), ig, image.ZP, draw.Src)
	return
}

//DrawImage 在当前画版上绘制图片
func (img *Image) DrawImage(p string, sx int, sy int, ex int, ey int) (err error) {
	m, err := decodeImage(p)
	if err != nil {
		return
	}
	draw.Draw(img.m, image.Rect(sx, sy, ex, ey), m, image.Pt(0, 0), draw.Over)
	return
}

//DrawFont 绘制字体
func (img *Image) DrawFont(fontPath string, text string, col string, fontSize float64, sx int, sy int) (err error) {
	data, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		return
	}

	c := freetype.NewContext()
	c.SetDst(img.m)
	c.SetClip(img.m.Bounds())
	switch col {
	case "000000":
		c.SetSrc(image.NewUniform(color.Gray16{0}))
	case "ffffff":
		c.SetSrc(image.NewUniform(color.Gray16{0xffff}))
	default:
		r, g, b, err := colorToRGB(col)
		if err != nil {
			return err
		}
		c.SetSrc(image.NewUniform(color.RGBA{R: r, G: g, B: b, A: 1}))
	}

	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetHinting(font.HintingNone)
	_, err = c.DrawString(text, freetype.Pt(sx, sy))
	return
}

//DrawImageWithScale 绘制图片并缩放原始图片
func (img *Image) DrawImageWithScale(p string, sx int, sy int, w int, h int) (err error) {
	m1, err := decodeImage(p)
	if err != nil {
		return
	}
	//缩放图片
	m := resize.Resize(uint(w), uint(h), m1, resize.Lanczos3)
	draw.Draw(img.m, image.Rect(sx, sy, sx+w, sy+h), m, image.Pt(0, 0), draw.Over)
	return
}

//Save 保存图片到指定路径
func (img *Image) Save(path string) error {
	imgfile, _ := os.Create(path)
	defer imgfile.Close()
	if strings.HasSuffix(path, ".jpg") {
		return jpeg.Encode(imgfile, img.m, &jpeg.Options{Quality: 70})
	}
	return png.Encode(imgfile, img.m)
}

func decodeImage(p string) (m image.Image, err error) {
	f1, err := os.Open(p)
	if err != nil {
		return
	}
	if strings.HasSuffix(p, ".jpg") {
		m, err = jpeg.Decode(f1)
		return
	} else if strings.HasSuffix(p, ".png") {
		m, err = png.Decode(f1)
		return
	}
	return nil, fmt.Errorf("图片格式不支持")

}

func colorToRGB(sc string) (red, green, blue uint8, err error) {
	color64, err := strconv.ParseInt(sc, 16, 32) //字串到数据整型
	if err != nil {
		return
	}
	c := int(color64) //类型强转
	red = uint8(c >> 16)
	green = uint8((c & 0x00FF00) >> 8)
	blue = uint8(c & 0x0000FF)
	return
}
