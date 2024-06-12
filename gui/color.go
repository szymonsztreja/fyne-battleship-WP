package gui

// // Color represents an RGB color.
// type Color struct {
// 	Red   uint8
// 	Green uint8
// 	Blue  uint8
// }

// // NewColor returns a new color. Parameters are red, green and blue values.
// func NewColor(r, g, b uint8) Color {
// 	return Color{Red: r, Green: g, Blue: b}
// }

// var (
// 	White = Color{Red: 208, Green: 208, Blue: 208}
// 	Black = Color{Red: 21, Green: 21, Blue: 21}
// 	Blue  = Color{Red: 108, Green: 153, Blue: 187}
// 	Red   = Color{Red: 172, Green: 65, Blue: 66}
// 	Grey  = Color{Red: 105, Green: 105, Blue: 105}
// 	Green = Color{Red: 126, Green: 142, Blue: 0}
// )

// Color represents a custom color type that implements color.Color
type Color struct {
	R, G, B, A uint8
}

// RGBA implements the color.Color interface for the Color type
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) * 0x101
	g = uint32(c.G) * 0x101
	b = uint32(c.B) * 0x101
	a = uint32(c.A) * 0x101
	return r, g, b, a
}

var (
	White = Color{R: 208, G: 208, B: 208, A: 255}
	Black = Color{R: 21, G: 21, B: 21, A: 255}
	Blue  = Color{R: 108, G: 153, B: 187, A: 255}
	Red   = Color{R: 172, G: 65, B: 66, A: 255}
	Grey  = Color{R: 105, G: 105, B: 105, A: 255}
	Green = Color{R: 126, G: 142, B: 0, A: 255}
)
