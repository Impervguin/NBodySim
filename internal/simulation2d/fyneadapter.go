package simulation2d

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type FyneDrawer struct {
	content *fyne.Container
}

func NewFyneDrawer(content *fyne.Container) *FyneDrawer {
	return &FyneDrawer{content: content}
}

func (d *FyneDrawer) DrawCircle(radius float64, color color.Color, xc, yc int64) {
	circle := canvas.NewCircle(color)
	circle.Move(fyne.NewPos(float32(xc)-float32(radius)/2, float32(yc)-float32(radius)/2))
	circle.Resize(fyne.NewSize(float32(radius), float32(radius)))
	circle.FillColor = color
	d.content.Add(circle)
}

func (d *FyneDrawer) DrawText(text string, color color.Color, x, y int64) {
	label := canvas.NewText(text, color)
	label.Move(fyne.NewPos(float32(x), float32(y)))
	d.content.Add(label)
}

func (d *FyneDrawer) Clear() {
	d.content.RemoveAll()
}

func (d *FyneDrawer) Refresh() {
	d.content.Refresh()
}
