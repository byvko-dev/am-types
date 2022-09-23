package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type PlayerSession struct {
	TotalBattles   int `bson:"totalBattles" json:"totalBattles"`
	LastBattleTime int `bson:"lastBattleTime" json:"lastBattleTime"`

	RegularBattles  statistics.StatsFrame                `bson:"regularBattles" json:"regularBattles"`
	RegularVehicles map[int]statistics.VehicleStatsFrame `bson:"regularVehicles" json:"regularVehicles"`

	RatingBattles statistics.StatsFrame `bson:"ratingBattles" json:"ratingBattles"`
}

// Adds playerSession b to playerSession a
func (a *PlayerSession) Add(b *PlayerSession) {
	a.TotalBattles += b.TotalBattles
	if a.LastBattleTime < b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}

	a.RegularBattles.Add(&b.RegularBattles)
	for id, bStats := range b.RegularVehicles {
		if aStats, ok := a.RegularVehicles[id]; ok {
			aStats.Add(&bStats)
		} else {
			a.RegularVehicles[id] = bStats
		}
	}

	a.RatingBattles.Add(&b.RatingBattles)
}

// Subtract playerSession b from playerSession a
func (a *PlayerSession) Subtract(b *PlayerSession) {
	a.TotalBattles -= b.TotalBattles
	if a.LastBattleTime > b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}

	a.RegularBattles.Subtract(&b.RegularBattles)
	for id, bStats := range b.RegularVehicles {
		if aStats, ok := a.RegularVehicles[id]; ok {
			aStats.Subtract(&bStats)
		}
	}

	a.RatingBattles.Subtract(&b.RatingBattles)
}
