package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type VehicleStats struct {
	statistics.VehicleStatsFrame `bson:",inline"`
	Achievements                 statistics.AchievementsFrame `json:"achievements" bson:"achievements"`
	Ratings                      map[string]int               `json:"ratings" bson:"ratings"`
	TankName                     map[string]string            `json:"tank_name" bson:"tank_name"`
	TankTier                     int                          `json:"tank_tier" bson:"tank_tier"`
}

func (v *VehicleStats) Add(b *VehicleStats) {
	v.VehicleStatsFrame.Add(&b.VehicleStatsFrame)
	v.Achievements.Add(&b.Achievements)
	for key, value := range b.Ratings {
		v.Ratings[key] += value
	}
}

func (v *VehicleStats) Subtract(b *VehicleStats) {
	v.VehicleStatsFrame.Subtract(&b.VehicleStatsFrame)
	v.Achievements.Subtract(&b.Achievements)
	for key, value := range b.Ratings {
		v.Ratings[key] -= value
	}
}
