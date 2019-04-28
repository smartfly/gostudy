package main

import (
	"fmt"
	"image"
)

/**
 * @author taohu
 * @date 19-4-28
 * @time 下午3:42
 * @desc 图像
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

/*
	package image

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle
		At(x, y int) color.Color
	}
*/

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
