package users

import (
	"image/color"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pin struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	URL         string             `bson:"url" json:"url"`
	Name        string             `bson:"name" json:"name"`
	Label       string             `bson:"label" json:"label"`
	Weight      int                `bson:"weight" json:"weight"`
	Hidden      bool               `bson:"hidden" json:"hidden"`
	Description string             `bson:"description" json:"description"`

	Size      int        `bson:"size" json:"size"`
	Glow      bool       `bson:"glow" json:"glow"`
	GlowColor color.RGBA `bson:"glowColor" json:"glowColor"`
}
