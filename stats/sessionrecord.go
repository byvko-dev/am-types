package stats

import (
	"sort"
	"time"

	"aftermath.link/repo/am-types/wargaming"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Used for conditional TTL indexes, determines how long this record will be stored for
type StoreDuration string

const (
	StoreDurationOneDay StoreDuration = "oneDay"
	StoreDuration7Days  StoreDuration = "sevenDays"
	StoreDuration30Days StoreDuration = "thirtyDays"
	StoreDuration60Days StoreDuration = "sixtyDays"
	StoreDuration90Days StoreDuration = "ninetyDays"
)

func (sd StoreDuration) IsValid() bool {
	switch sd {
	case StoreDurationOneDay, StoreDuration7Days, StoreDuration30Days, StoreDuration60Days, StoreDuration90Days:
		return true
	default:
		return false
	}
}

// Represents a record in the stats collection
type SessionRecord struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"` // Unique Record ID
	PlayerID wargaming.PlayerID `bson:"playerId" json:"playerId"`

	CreationTime   time.Time `bson:"creationTime" json:"creationTime"`     // Time of record creation
	LastUpdateTime time.Time `bson:"lastUpdateTime" json:"lastUpdateTime"` // Time of last record update

	// Stats
	ShardedStats ShardedStats  `bson:"shardedStats" json:"shardedStats"`
	CurrentStats PlayerSession `bson:"currentStats" json:"currentStats"`

	// Store duration
	StoreDuration `bson:"storeDuration" json:"storeDuration"`
}

type ShardedStats struct {
	TotalShards  int                 `bson:"totalShards" json:"totalShards"`
	InitialStats PlayerSession       `bson:"initialStats" json:"initialStats"`
	Progression  []ProgressionRecord `bson:"progression" json:"progression"`
}

type ProgressionRecord struct {
	Timestamp  time.Time     `bson:"timestamp" json:"timestamp"`
	StatsShard PlayerSession `bson:"stats" json:"stats"`
}

// Sorts the progression records by timestamp
func (s *ShardedStats) SortProgression() {
	sort.Slice(s.Progression, func(i, j int) bool {
		return s.Progression[i].Timestamp.Before(s.Progression[j].Timestamp)
	})
}

// Checks if progression records are sorted
func (s *ShardedStats) IsProgressionSorted() bool {
	if len(s.Progression) <= 1 {
		return true
	}
	for i := 1; i < len(s.Progression); i++ {
		if s.Progression[i].Timestamp.Before(s.Progression[i-1].Timestamp) {
			return false
		}
	}
	return true
}

// Adds a new progression record to the sharded stats and sorts the progression
func (s *ShardedStats) AddProgressionRecord(record ProgressionRecord) {
	s.Progression = append(s.Progression, record)
	s.SortProgression()
}

// Computes the progression of the stats record and returns the stats frame
func (s *ShardedStats) ComputeStatsFrame(maximumShards int) (stats PlayerSession) {
	if maximumShards > s.TotalShards {
		maximumShards = s.TotalShards
	}

	stats = s.InitialStats
	for i := 0; i < maximumShards; i++ {
		stats.Add(&s.Progression[i].StatsShard)
	}
	return stats
}
