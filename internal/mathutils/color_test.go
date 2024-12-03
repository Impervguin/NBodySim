package mathutils

import (
	"image/color"
	"testing"
)

func TestMultRGBA64(t *testing.T) {
	c := color.RGBA{R: 255, G: 128, B: 0, A: 255}
	coeff := 0.5
	result := MultRGBA64(c, coeff)
	if result.R != 32767 || result.G != 16448 || result.B != 0 || result.A != 65535 {
		t.Error("Expected (128, 64, 0, 128), got", result)
		t.Error(c.RGBA())
	}
}

func TestMultRGBA64AboveMax(t *testing.T) {
	c := color.RGBA{R: 255, G: 128, B: 0, A: 255}
	coeff := 2.0
	result := MultRGBA64(c, coeff)
	if result.R != 65535 || result.G != 65535 || result.B != 0 || result.A != 65535 {
		t.Error("Expected (65535, 65535, 0, 65535), got", result)
	}
}

func TestToRGBA64(t *testing.T) {
	c := color.RGBA{R: 255, G: 128, B: 0, A: 255}
	result := ToRGBA64(c)
	if result.R != 65535 || result.G != 32896 || result.B != 0 || result.A != 65535 {
		t.Error("Expected (255, 128, 0, 255), got", result)
	}
}

func TestToRGBA(t *testing.T) {
	c := color.RGBA{R: 255, G: 128, B: 0, A: 255}
	result := ToRGBA(c)
	if result.R != 255 || result.G != 128 || result.B != 0 || result.A != 255 {
		t.Error("Expected (127, 64, 0, 255), got", result)
	}
}

func TestAddRGBA(t *testing.T) {
	c1 := color.RGBA{R: 255, G: 128, B: 0, A: 255}
	c2 := color.RGBA{R: 128, G: 64, B: 0, A: 128}
	result := AddRGBA(c1, c2)
	if result.R != 255 || result.G != 192 || result.B != 0 || result.A != 255 {
		t.Error("Expected (255, 192, 0, 383), got", result)
	}
}

func TestAddRGBA64(t *testing.T) {
	c1 := color.RGBA64{R: 65535, G: 32896, B: 0, A: 65535}
	c2 := color.RGBA64{R: 32767, G: 16448, B: 0, A: 32767}
	result := AddRGBA64(c1, c2)
	if result.R != 65535 || result.G != 49344 || result.B != 0 || result.A != 65535 {
		t.Error("Expected (65535, 49344, 0, 65535), got", result)
	}
}

func TestLinearColorInterpolation(t *testing.T) {
	colors := LinearColorInterpolation(0, 25, color.RGBA64{250, 250, 250, 250}, color.RGBA64{0, 0, 0, 0})
	if len(colors) != 26 {
		t.Error("Expected 26 colors, got", len(colors))
	}
	dc := 250. / 25.
	c := 250.

	for i := 0; i < 26; i++ {
		r, g, b, a := colors[i].RGBA()
		cint := ToInt(c)
		if r != uint32(cint) || g != uint32(cint) || b != uint32(cint) || a != uint32(cint) {
			t.Errorf("Expected (%v, %v, %v, %v), got %v", cint, cint, cint, cint, colors[i])
		}
		c -= dc
	}
}

func TestLinearColorInterpolationSwap(t *testing.T) {
	colors := LinearColorInterpolation(25, 0, color.RGBA64{250, 250, 250, 250}, color.RGBA64{0, 0, 0, 0})
	if len(colors) != 26 {
		t.Error("Expected 26 colors, got", len(colors))
	}
	dc := 250. / 25.
	c := 0.

	for i := 0; i < 26; i++ {
		r, g, b, a := colors[i].RGBA()
		cint := ToInt(c)
		if r != uint32(cint) || g != uint32(cint) || b != uint32(cint) || a != uint32(cint) {
			t.Errorf("Expected (%v, %v, %v, %v), got %v", cint, cint, cint, cint, colors[i])
		}
		c += dc
	}
}
