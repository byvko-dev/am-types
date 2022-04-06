package style

import (
	"image/color"
	"reflect"
)

// Block content alignment
type AlignItemsValue string

const AlignItemsVertical AlignItemsValue = "vertical"
const AlignItemsHorizontal AlignItemsValue = "horizontal"

type Style struct {
	Color           color.RGBA `json:"color" bson:"color"`
	BackgroundColor color.RGBA `json:"backgroundColor" bson:"backgroundColor"`
	BackgroundImage []byte     `json:"backgroundImage" bson:"backgroundImage"`

	// Font
	FontName string  `json:"fontFamily" bson:"fontFamily"`
	FontSize float64 `json:"fontSize" bson:"fontSize"`

	// Border
	BorderRadius float64 `json:"borderRadius" bson:"borderRadius"`

	// Alignment
	AlignItems AlignItemsValue `json:"alignItems" bson:"alignItems"`
	Gap        float64         `json:"gap" bson:"gap"`

	// Spacing / Sizing
	PaddingX float64 `json:"paddingX" bson:"paddingX"`
	PaddingY float64 `json:"paddingY" bson:"paddingY"`
	Grow     bool    `json:"grow" bson:"grow"`
}

// Merge o into s, overwriting s values with o values if they are not nil
func (s Style) Merge(o Style) Style {
	if reflect.TypeOf(o) != reflect.TypeOf(s) {
		return s
	}
	var result Style
	for i := 0; i < reflect.TypeOf(o).NumField(); i++ {
		// handler slices
		if reflect.TypeOf(o).Field(i).Type.Kind() == reflect.Slice {
			if !reflect.ValueOf(o).Field(i).IsNil() {
				reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(o).Field(i))
			} else {
				reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(s).Field(i))
			}
			continue
		}

		if reflect.ValueOf(o).Field(i).Interface() != reflect.Zero(reflect.TypeOf(o).Field(i).Type).Interface() {
			reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(o).Field(i))
		} else {
			reflect.ValueOf(&result).Elem().Field(i).Set(reflect.ValueOf(s).Field(i))
		}
	}
	return result
}
