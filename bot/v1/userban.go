package bot

type UserBan struct {
	Banned    bool   `json:"banned"`
	BanReason string `json:"banReason"`
	BanExpiry int    `json:"banExpiry"`
}
