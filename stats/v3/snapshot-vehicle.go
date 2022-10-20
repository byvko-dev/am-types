package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type SnapshotVehicleStats struct {
	statistics.VehicleStatsFrame `bson:",inline"`
	Achievements                 statistics.AchievementsFrame `json:"achievements" bson:"achievements,omitempty"`
	Ratings                      map[string]int               `json:"ratings" bson:"ratings,omitempty"`
}
