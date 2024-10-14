package simulation2d

import "image/color"

type Sim2dDrawer interface {
	// Draws the circle with (xc, yc) center
	DrawCircle(radius float64, color color.Color, xc, yc int64)
	// Draws the text at (x, y)(left top rectangleangle)
	DrawText(text string, color color.Color, x, y int64)
	// Clears the entire canvas
	Clear()
	// Updates screen
	Refresh()
}
