package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type Frame struct {
	Total        statistics.StatsFrame        `json:"total" bson:"total"`
	Achievements statistics.AchievementsFrame `json:"achievements" bson:"achievements"`
	Ratings      map[string]int               `json:"ratings" bson:"ratings"`
	Vehicles     map[int]SessionVehicleStats  `json:"vehicles" bson:"vehicles"`
}

type CompleteFrame struct {
	Regular Frame `json:"regular" bson:"regular"`
	Rating  Frame `json:"rating" bson:"rating"`
}
