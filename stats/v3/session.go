package stats

import "github.com/byvko-dev/am-types/wargaming/v2/statistics"

type Session struct {
	AccountID      int64  `json:"account_id" bson:"account_id"`
	AccountName    string `json:"account_name" bson:"account_name"`
	AccountClanTag string `json:"account_clan_tag" bson:"account_clan_tag"`

	TotalBattles   int `json:"totalBattles" bson:"totalBattles"`
	LastBattleTime int `json:"lastBattleTime" bson:"lastBattleTime"`

	Stats CompleteFrame `json:"stats" bson:"stats"`
}

type VehicleStats struct {
	statistics.VehicleStatsFrame `bson:",inline"`
	Achievements                 statistics.AchievementsFrame `json:"achievements" bson:"achievements"`
	Ratings                      map[string]int               `json:"ratings" bson:"ratings"`
	TankName                     map[string]string            `json:"tank_name" bson:"tank_name"`
	TankTier                     int                          `json:"tank_tier" bson:"tank_tier"`
}
