package settings

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenerationSettings struct {
	ID               primitive.ObjectID `json:"id" firestore:"id" bson:"_id,omitempty"`
	Style            string             `json:"style" firestore:"style" bson:"style"`
	Preset           string             `json:"preset" firestore:"preset" bson:"preset"`
	Options          Options            `json:"options,omitempty" firestore:"options" bson:"options"`
	UseCustomOptions bool               `json:"useCustomOptions,omitempty" firestore:"useCustomOptions" bson:"useCustomOptions"`
	LastUsed         time.Time          `json:"lastUsed,omitempty" firestore:"lastUsed" bson:"lastUsed"`

	Player struct {
		ID    int    `json:"id,omitempty" firestore:"id" bson:"id"`
		Realm string `json:"realm,omitempty" firestore:"realm" bson:"realm"`
	} `json:"player,omitempty" firestore:"player" bson:"player"`

	OwnerId string `json:"ownerId,omitempty" firestore:"ownerId" bson:"ownerId"`
	Locale  string `json:"locale,omitempty" firestore:"locale" bson:"locale"`
}

func (s *GenerationSettings) Validate() error {
	if s.Player.ID == 0 {
		return errors.New("player ID is not set")
	}
	if !CheckRealm(s.Player.Realm) {
		return errors.New("realm is not valid")
	}
	var blank Options
	if fmt.Sprintf("%v", s.Options) == fmt.Sprintf("%v", blank) {
		return errors.New("options are not set")
	}
	return nil
}

func CheckRealm(realm string) bool {
	switch realm {
	case "NA":
		return true
	case "EU":
		return true
	case "RU":
		return true
	case "ASIA":
		return true
	default:
		return false
	}

}
