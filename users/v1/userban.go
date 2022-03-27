package users

type UserBan struct {
	Banned    bool   `json:"banned" bson:"banned"`
	BanReason string `json:"banReason" bson:"banReason"`
	BanExpiry int    `json:"banExpiry" bson:"banExpiry"`
}
