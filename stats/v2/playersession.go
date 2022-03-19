package stats

import "github.com/byvko-dev/am-types/wargaming"

type PlayerSession struct {
	TotalBattles   int `bson:"totalBattles" json:"totalBattles"`
	LastBattleTime int `bson:"lastBattleTime" json:"lastBattleTime"`

	RegularBattles  wargaming.StatsFrame                `bson:"regularBattles" json:"regularBattles"`
	RegularVehicles map[int]wargaming.VehicleStatsFrame `bson:"regularVehicles" json:"regularVehicles"`

	RatingBattles wargaming.StatsFrame `bson:"ratingBattles" json:"ratingBattles"`
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

// Substract playerSession b from playerSession a
func (a *PlayerSession) Substract(b *PlayerSession) {
	a.TotalBattles -= b.TotalBattles
	if a.LastBattleTime > b.LastBattleTime {
		a.LastBattleTime = b.LastBattleTime
	}

	a.RegularBattles.Substract(&b.RegularBattles)
	for id, bStats := range b.RegularVehicles {
		if aStats, ok := a.RegularVehicles[id]; ok {
			aStats.Substract(&bStats)
		}
	}

	a.RatingBattles.Substract(&b.RatingBattles)
}
