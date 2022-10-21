package users

import (
	"image/color"
)

type PinType string

const PinTypeRemoteImage PinType = "remoteImage"
const PinTypeIcon PinType = "icon"

type Pin struct {
	ID          string `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string `bson:"name" json:"name"`
	Label       string `bson:"label" json:"label"`
	Description string `bson:"description" json:"description"`

	Size   int     `bson:"size" json:"size"`
	Type   PinType `bson:"type" json:"type"`
	Weight int     `bson:"weight" json:"weight"`
	Hidden bool    `bson:"hidden" json:"hidden"`

	Glow       bool       `bson:"glow" json:"glow"`
	GlowColor  color.RGBA `bson:"glowColor" json:"glowColor"`
	Content    string     `bson:"content" json:"content"`
	Attributes []string   `bson:"attributes" json:"attributes"`
}
