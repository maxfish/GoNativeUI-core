package utils

type Color [4]float32

var TransparentColor = Color{0,0,0,0}

// NewColor creates a new color from float components (0->1)
func NewColor(r, g, b, a float32) Color {
	return Color{r, g, b, a}
}

// NewColorInt creates a new color from 8bit components (0->255)
func NewColorInt(r, g, b, a uint8) Color {
	return Color{
		float32(r) / 255.0,
		float32(g) / 255.0,
		float32(b) / 255.0,
		float32(a) / 255.0,
	}
}

// NewColorHex creates a new color from an hex number
func NewColorHex(hex uint32) Color {
	c := Color{}
	c[3] = float32(hex&0xFF) / 255
	hex >>= 8
	c[2] = float32(hex&0xFF) / 255
	hex >>= 8
	c[1] = float32(hex&0xFF) / 255
	hex >>= 8
	c[0] = float32(hex&0xFF) / 255
	return c
}

// NewColorGrayInt creates a gray shade from float components (0->1)
func NewColorGrayInt(g, a uint8) Color {
	return Color{
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(a) / 255.0,
	}
}

func (c Color) Scaled(scale float32) Color {
	return Color{
		c[0] * scale,
		c[1] * scale,
		c[2] * scale,
		c[3] * scale,
	}
}
