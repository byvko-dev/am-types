package stats

type ClanSnapshot struct {
	ClanID    int64 `json:"account_id" bson:"account_id"`
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	IsManual  bool  `json:"is_manual" bson:"is_manual"`

	TotalMembers int   `json:"total_members" bson:"total_members"`
	MemberIDs    []int `json:"member_ids" bson:"member_ids"`

	Stats CompleteFrame `json:"stats" bson:"stats"`
}
