package stats

import "github.com/byvko-dev/am-types/wargaming/v1/statistics"

type PlayerDetails struct {
	playerClanData `json:"clan"`
	ID             int            `json:"account_id"`
	Name           string         `json:"nickname"`
	LastBattle     int            `json:"last_battle_time"`
	Stats          playerStatsRes `json:"statistics"`
	CareerWN8      int            `json:"career_wn8"`
	Realm          string         `json:"realm"`
}

// Player stats response
type playerStatsRes struct {
	Rating statistics.StatsFrame `json:"rating"`
	All    statistics.StatsFrame `json:"all"`
}

// PlayerProfile -
type playerClanData struct {
	ClanTag  string `json:"tag,omitempty"`
	ClanName string `json:"name,omitempty"`
	ClanID   int    `json:"clan_id,omitempty"`
}
