package discord

// Full Discord Profile and auth token data per https://discord.com/developers/docs/resources/user#user-object
type DiscordProfile struct {
	Token AccessToken `json:"token" firestore:"token" bson:"token"`

	ID          string `json:"id" firestore:"id" bson:"id"`
	Username    string `json:"username" firestore:"username" bson:"username"`
	Discrim     string `json:"discriminator" firestore:"discriminator" bson:"discriminator"`
	Avatar      string `json:"avatar" firestore:"avatar" bson:"avatar"`
	Mfa         bool   `json:"mfa_enabled" firestore:"mfa_enabled" bson:"mfa_enabled"`
	Banner      string `json:"banner" firestore:"banner" bson:"banner"`
	AccentColor int    `json:"accent_color" firestore:"accent_color" bson:"accent_color"`
	Locale      string `json:"locale" firestore:"locale" bson:"locale"`
	Verified    bool   `json:"verified" firestore:"verified" bson:"verified"`
	Email       string `json:"email" firestore:"email" bson:"email"`
	Flags       int    `json:"flags" firestore:"flags" bson:"flags"`
	PremiumType int    `json:"premium_type" firestore:"premium_type" bson:"premium_type"`
	PublicFlags int    `json:"public_flags" firestore:"public_flags" bson:"public_flags"`
}
