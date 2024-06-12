package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

const (
	fieldWidth  = 20
	fieldHeight = 1
)

var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

type tile struct {
	rec     *canvas.Rectangle
	txt     *canvas.Text
	btn     *stateButton
	widgets *fyne.Container
}

// Board represents a single board.
type Board struct {
	id    uuid.UUID
	cfg   *BoardConfig
	tiles []*fyne.Container
	ch    chan string

	x int
	y int
}

// BoardConfig holds configuration parameters for Board struct.
type BoardConfig struct {
	RulerColor Color
	TextColor  Color
	EmptyColor Color
	HitColor   Color
	MissColor  Color
	ShipColor  Color
	EmptyChar  byte
	HitChar    byte
	MissChar   byte
	ShipChar   byte
}

// NewBoardConfig returns a new config with default values.
func NewBoardConfig() *BoardConfig {
	return &BoardConfig{
		RulerColor: White,
		TextColor:  Black,
		EmptyColor: Blue,
		HitColor:   Red,
		MissColor:  Grey,
		ShipColor:  Green,
		EmptyChar:  '~',
		HitChar:    'H',
		MissChar:   'M',
		ShipChar:   'S',
	}
}

func (c *BoardConfig) getColor(state State) Color {
	switch state {
	case Hit:
		return c.HitColor
	case Miss:
		return c.MissColor
	case Ship:
		return c.ShipColor
	default:
		return c.EmptyColor
	}
}

func (c *BoardConfig) getChar(state State) byte {
	switch state {
	case Hit:
		return c.HitChar
	case Miss:
		return c.MissChar
	case Ship:
		return c.ShipChar
	default:
		return c.EmptyChar
	}
}

// func (b *Board) Display() *fyne.Container {
// 	// Create a new container for the board tiles
// 	grid := container.New(layout.NewGridLayout(11))

// 	for _, t := range b.tiles {
// 		if t == nil {
// 			continue
// 		}
// 		container := container.NewBorder(nil, nil, nil, nil, t.rec, t.txt)
// 		if t.btn != nil {
// 			container.Add(t.btn.button)
// 		}
// 		grid.Add(container)
// 	}

// 	return grid
// }

// NewBoard returns a new Board struct.
// X and Y are the coordinates of the top left corner of the board.
// If no config is provided, default values are used.
// func NewBoard(x, y int, cfg *BoardConfig) *Board {
// 	if cfg == nil {
// 		cfg = NewBoardConfig()
// 	}

// 	// grid := container.New(layout.NewGridLayout(11))

// 	b := &Board{
// 		id:  uuid.New(),
// 		cfg: cfg,
// 		ch:  make(chan string),
// 		x:   x,
// 		y:   y,
// 	}

// 	b.tiles = make([]*fyne.Container, 11*11)

// 	// Set rectangle and text
// 	for n := 1; n <= 10; n++ {
// 		newX := n * fieldWidth
// 		horizontal := &tile{
// 			rec: newRectangle(x+newX, y, fieldWidth, fieldHeight, cfg.RulerColor),
// 			txt: canvas.NewText(letters[n-1], cfg.TextColor),
// 		}
// 		horizontal.txt.Move(fyne.NewPos(float32(x+newX+fieldWidth/4), float32(y+fieldHeight/4)))
// 		b.tiles[n] = horizontal

// 		newY := n * fieldHeight
// 		vertical := &tile{
// 			rec: newRectangle(x, y+newY, fieldWidth, fieldHeight, cfg.RulerColor),
// 			txt: canvas.NewText(fmt.Sprintf("%d", n), cfg.TextColor),
// 		}
// 		vertical.txt.Move(fyne.NewPos(float32(x+fieldWidth/4), float32(y+newY+fieldHeight/4)))
// 		b.tiles[n*11] = vertical
// 	}

// 	// Here set container and a button
// 	// Creating buttons
// 	for y := 1; y <= 10; y++ {
// 		for x := 1; x <= 10; x++ {
// 			newX := x*fieldWidth + x
// 			newY := y*fieldHeight + y
// 			t := &tile{
// 				btn: newStateButton(fmt.Sprintf("%s%d", letters[x-1], y), fmt.Sprintf("%s%d", letters[x-1], y), b.ch),
// 			}
// 			t.btn.button.Move(fyne.NewPos(float32(newX), float32(newY)))
// 			t.widgets = container.NewStack(t.btn.button)
// 			b.tiles[x+y*11] = t
// 		}
// 	}
// 	return b
// }

// NewBoard returns a new Board struct.
// X and Y are the coordinates of the top left corner of the board.
// If no config is provided, default values are used.
func NewBoard(x, y int, cfg *BoardConfig) *Board {
	if cfg == nil {
		cfg = NewBoardConfig()
	}

	b := &Board{
		id:  uuid.New(),
		cfg: cfg,
		ch:  make(chan string),
		x:   x,
		y:   y,
	}

	b.tiles = make([]*fyne.Container, 11*11)
	cellSize := fyne.NewSize(50, 50) // Set the fixed size for each cell

	// Set rectangle and text
	for n := 0; n <= 10; n++ {
		var horizontal *fyne.Container
		if n == 0 {
			horizontal = container.NewWithoutLayout(
				canvas.NewRectangle(cfg.RulerColor),
				canvas.NewText("", cfg.ShipColor),
			)
		} else {
			horizontal = container.NewWithoutLayout(
				canvas.NewRectangle(cfg.RulerColor),
				canvas.NewText(letters[n-1], cfg.ShipColor),
			)
		}
		horizontal.Objects[1].Move(fyne.NewPos(cellSize.Height/2, cellSize.Width/2)) // Adjust position as needed
		b.tiles[n] = container.New(layout.NewGridWrapLayout(cellSize), horizontal)

		var vertical *fyne.Container
		if n == 0 {
			vertical = container.NewWithoutLayout(
				canvas.NewRectangle(cfg.RulerColor),
				canvas.NewText("", cfg.ShipColor),
			)
		} else {
			vertical = container.NewWithoutLayout(
				canvas.NewRectangle(cfg.RulerColor),
				canvas.NewText(letters[n-1], cfg.MissColor),
			)
		}
		horizontal.Objects[0].Move(fyne.NewPos(cellSize.Height/2, cellSize.Width/2))
		b.tiles[n*11] = container.New(layout.NewGridWrapLayout(cellSize), vertical)
	}

	// Creating buttons
	for y := 1; y <= 10; y++ {
		for x := 1; x <= 10; x++ {
			t := container.New(layout.NewGridWrapLayout(cellSize),
				widget.NewButton(fmt.Sprintf("%s%d", letters[x-1], y), func() {}))
			b.tiles[x+y*11] = t
		}
	}

	return b
}

func (b *Board) GetBoardTiles() []fyne.CanvasObject {
	var canvasObjects []fyne.CanvasObject
	for _, tile := range b.tiles {
		if tile != nil {
			canvasObjects = append(canvasObjects, tile)
		}
	}
	return canvasObjects
}

// SetStates sets the states of the board. The states are represented
// as a 10x10 matrix, where the first index is the X coordinate and
// the second index is the Y coordinate.
// Example: states[0][0] is the state of the field A1.
// func (b *Board) SetStates(states [10][10]State) {
// 	for y := 1; y <= 10; y++ {
// 		for x := 1; x <= 10; x++ {
// 			state := states[x-1][y-1]
// 			color := b.cfg.getColor(state)
// 			b.tiles[x+y*11].rec.FillColor = color
// 			b.tiles[x+y*11].txt.Color = color
// 			b.tiles[x+y*11].txt.Text = string(b.cfg.getChar(state))
// 		}
// 	}
// }
