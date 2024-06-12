package main

import (
	"fyne-battleship-WP/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	win := myApp.NewWindow("Board Test")

	// // Create a new Board
	board := gui.NewBoard(50, 50, nil)

	tiles := board.GetBoardTiles()
	grid := container.New(layout.NewGridLayout(11), tiles...)

	win.SetContent(grid)
	win.Resize(fyne.NewSize(600, 600))
	win.ShowAndRun()
}
