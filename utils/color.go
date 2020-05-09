package utils

type ColorF [4]float32

// ColorRGBA creates a new color from float components (0->1)
func ColorRGBA(r, g, b, a float32) ColorF {
	return ColorF{r, g, b, a}
}

// ColorRGBAi creates a new color from 8bit components (0->255)
func ColorRGBAi(r, g, b, a uint8) ColorF {
	return ColorF{
		float32(r) / 255.0,
		float32(g) / 255.0,
		float32(b) / 255.0,
		float32(a) / 255.0,
	}
}

// ColorGrayi creates a gray shade from float components (0->1)
func ColorGrayi(g, a uint8) ColorF {
	return ColorF{
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(g) / 255.0,
		float32(a) / 255.0,
	}
}

func ColorScaled(c ColorF, scale float32) ColorF {
	return ColorF{
		c[0] * scale,
		c[1] * scale,
		c[2] * scale,
		c[3] * scale,
	}
}
