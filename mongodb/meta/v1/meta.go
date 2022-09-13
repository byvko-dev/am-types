package meta

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Meta struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (m *Meta) MarshalBSON() ([]byte, error) {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	m.UpdatedAt = time.Now()
	type my Meta
	return bson.Marshal((*my)(m))
}
