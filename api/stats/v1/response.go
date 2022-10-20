package stats

import (
	"time"

	"github.com/byvko-dev/am-types/stats/v3"
)

type ResponsePayload struct {
	AccountID int       `json:"account_id" bson:"account_id"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`

	Account  stats.AccountInfo   `json:"account" bson:"account"`
	Session  stats.CompleteFrame `json:"session" bson:"session"`
	Snapshot stats.CompleteFrame `json:"snapshot" bson:"snapshot"`
}
