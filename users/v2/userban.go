package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBan struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string             `json:"user_id" firestore:"user_id" bson:"user_id"`

	Lifted       bool   `json:"lifted" firestore:"lifted" bson:"lifted"`
	LiftedReason string `json:"lifted_reason" firestore:"lifted_reason" bson:"lifted_reason"`

	Reason     string    `json:"reason" bson:"reason"`
	Notified   bool      `json:"notified" bson:"notified"`
	Expiration time.Time `json:"expiration" bson:"expiration"`
}
