package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type AccountSnapshot struct {
	AccountID int64 `json:"account_id" bson:"account_id"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	IsManual  bool  `json:"is_manual" bson:"is_manual"`

	TotalBattles   int `json:"totalBattles" bson:"totalBattles"`
	LastBattleTime int `json:"lastBattleTime" bson:"lastBattleTime"`

	Stats struct {
		Regular SnapshotStats `json:"regular" bson:"regular"`
		Rating  SnapshotStats `json:"rating" bson:"rating"`
	} `json:"stats" bson:"stats"`
}

type SnapshotStats struct {
	statistics.StatsFrame
	Vehicles map[int]statistics.VehicleStatsFrame `json:"vehicles" bson:"vehicles"`
}
