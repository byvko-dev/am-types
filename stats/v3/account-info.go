package stats

type AccountInfo struct {
	AccountID      int         `json:"account_id" bson:"account_id"`
	Nickname       string      `json:"nickname" bson:"nickname"`
	Realm          string      `json:"realm" bson:"realm"`
	Clan           AccountClan `json:"clan,omitempty" bson:"clan,omitempty"`
	LastBattleTime int         `json:"last_battle_time" bson:"last_battle_time"`
}

type AccountClan struct {
	ID       int    `json:"id" bson:"id"`
	Tag      string `json:"tag" bson:"tag"`
	Name     string `json:"name" bson:"name"`
	Role     string `json:"role" bson:"role"`
	JoinedAt int    `json:"joined_at" bson:"joined_at"`
	EmblemID int    `json:"emblem_id" bson:"emblem_id"`
}
