package stats

type AccountSnapshot struct {
	AccountID int64 `json:"account_id" bson:"account_id"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	IsManual  bool  `json:"is_manual" bson:"is_manual"`

	TotalBattles   int `json:"total_battles" bson:"total_battles"`
	LastBattleTime int `json:"last_battle_time" bson:"last_battle_time"`

	Stats CompleteFrame `json:"stats" bson:"stats"`
}
