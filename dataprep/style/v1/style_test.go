package style

import (
	"image/color"
	"testing"
)

func TestStyleOverwriteNonNil(t *testing.T) {
	t.Run("Merge", func(t *testing.T) {
		style := Style{
			Color:           color.RGBA{R: 255, G: 255, B: 255, A: 255},
			BackgroundColor: color.RGBA{R: 0, G: 0, B: 0, A: 0},

			// Font
			FontName:        "somefont",
			FontSize:        999,
			BackgroundImage: []byte{1, 2, 3},

			// Border
			BorderRadius: 12,

			// Alignment
			AlignItems: AlignItemsHorizontal,
			Gap:        2,
		}

		style2 := Style{
			Color:           color.RGBA{R: 0, G: 1, B: 0, A: 255},
			BackgroundImage: []byte{3, 2, 1},
			FontName:        "default",
			FontSize:        14,
			BorderRadius:    0, // nil, should not be overwritten
		}

		style = style.Merge(style2)

		if style.BackgroundImage[0] != 3 {
			t.Errorf("Expected BackgroundImage[0] to be 3, got %d", style.BackgroundImage[0])
		}
		if style.Color.R != 0 {
			t.Errorf("Expected style.Color.R to be 0, got %d", style.Color.R)
		}
		if style.FontName != "default" {
			t.Errorf("Expected style.FontName to be default, got %s", style.FontName)
		}
		if style.FontSize != 14 {
			t.Errorf("Expected style.FontSize to be 14, got %f", style.FontSize)
		}
		if style.BorderRadius != 12 {
			t.Errorf("Expected style.BorderRadius to be 0, got %f", style.BorderRadius)
		}
	})
}
