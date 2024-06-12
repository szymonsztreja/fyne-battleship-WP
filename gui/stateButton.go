package gui

import (
	"fyne.io/fyne/v2/widget"
)

type stateButton struct {
	button *widget.Button
	coord  string
	ch     chan<- string
}

// newClickableRectangle creates a new clickableRectangle
func newStateButton(label, coord string, ch chan<- string) *stateButton {
	cr := &stateButton{
		button: widget.NewButton(label, nil),
		coord:  coord,
		ch:     ch,
	}

	cr.button.OnTapped = cr.processClick

	return cr
}

// processClick processes the click event and sends coordinates through the channel
func (c *stateButton) processClick() {
	select {
	case c.ch <- c.coord:
	default:
		// drop if the channel is not ready to receive
	}
}
