package stats

import (
	"github.com/byvko-dev/am-types/mongodb/meta/v1"
	"github.com/byvko-dev/am-types/wargaming/v2/accounts"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents a record in the stats collection
type SessionRecord struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"` // Unique Record ID
	PlayerID accounts.PlayerID  `bson:"playerId" json:"playerId"`

	meta.Meta // CreatedAt, UpdatedAt

	// Stats
	PlayerSession `bson:"stats" json:"stats"`
}
