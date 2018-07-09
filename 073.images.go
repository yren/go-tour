/*
    image 包定义了 Image 接口:
    package image
    type Image interface {
        ColorModel() color.ColorModel
        Bounds() Reactangle
        At(x, y int) color.Color
    }
*/

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
