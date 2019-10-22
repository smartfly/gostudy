package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 下午3:48
 * @desc 练习：图像
 * To change this template use File | Settings | File Templates | Includes | File Header
**/
type Image struct{}

// 颜色模式
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// 图片边界
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

// 某个点的颜色
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
