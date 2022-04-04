package block

type Style struct {
	Color string `json:"color" bson:"color"`

	AlignItems AlignItemsValue `json:"alignItems" bson:"alignItems"`
	Grow       bool            `json:"grow" bson:"grow"`
	Gap        int             `json:"gap" bson:"gap"`

	Padding int `json:"padding" bson:"padding"`
	Margin  int `json:"margin" bson:"margin"`
}

// Block content alignment
type AlignItemsValue string

const AlignItemsVertical AlignItemsValue = "vertical"
const AlignItemsHorizontal AlignItemsValue = "horizontal"
