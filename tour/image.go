package main

import (
	image2 "image"
	"image/color"
)

func Pic(dx, dy int) [][]uint8 {
	var u = make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		u[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			u[i][j] = uint8(i * j)
		}
	}
	return u
}

const length = 200

const width = 100

type Image struct{}

func (image Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (image Image) Bounds() image2.Rectangle {
	return image2.Rectangle{Min: image2.Point{}, Max: image2.Point{X: length, Y: width}}
}

func (image Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x - y),
		G: uint8(x * y),
		B: uint8(x + y),
		A: 255}
}
