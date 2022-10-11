package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type SnapshotStats struct {
	Total        statistics.StatsFrame        `json:"total" bson:"total"`
	Achievements statistics.AchievementsFrame `json:"achievements" bson:"achievements,omitempty"`
	Vehicles     map[int]SnapshotVehicleStats `json:"vehicles" bson:"vehicles,omitempty"`
}

type SnapshotVehicleStats struct {
	statistics.VehicleStatsFrame `bson:",inline"`
	Achievements                 statistics.AchievementsFrame `json:"achievements" bson:"achievements,omitempty"`
}
