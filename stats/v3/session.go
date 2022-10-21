package stats

type Session struct {
	AccountID      int64  `json:"account_id" bson:"account_id"`
	AccountName    string `json:"account_name" bson:"account_name"`
	AccountClanTag string `json:"account_clan_tag" bson:"account_clan_tag"`

	TotalBattles   int `json:"totalBattles" bson:"totalBattles"`
	LastBattleTime int `json:"lastBattleTime" bson:"lastBattleTime"`

	Stats CompleteFrame `json:"stats" bson:"stats"`
}
