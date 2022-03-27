package discord

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type" firestore:"grant_type" bson:"grant_type"`
	ClientSecret string `json:"client_secret" firestore:"client_secret" bson:"client_secret"`
	ClientID     string `json:"client_id" firestore:"client_id" bson:"client_id"`
	RefreshToken string `json:"refresh_token" firestore:"refresh_token" bson:"refresh_token"`
}
