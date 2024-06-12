package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type rectangle struct {
	*canvas.Rectangle
	coord string
	ch    chan<- string
}

func newRectangle(x, y, width, height int, color color.Color) *canvas.Rectangle {
	rec := canvas.NewRectangle(color)
	rec.Move(fyne.NewPos(float32(x), float32(y)))
	rec.Resize(fyne.NewSize(float32(width), float32(height)))
	return rec
}
