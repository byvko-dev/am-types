package users

type ExternalService string

const ExternalServiceWargaming ExternalService = "wargaming"
const ExternalServiceDiscord ExternalService = "discord"

type ExternalProfile struct {
	ExternalProfileID
	AccessToken  string            `json:"access_token" firestore:"access_token" bson:"access_token"`
	RefreshToken string            `json:"refresh_token" firestore:"refresh_token" bson:"refresh_token"`
	Meta         map[string]string `json:"meta" firestore:"meta" bson:"meta"`
}

type ExternalProfileID struct {
	Service    ExternalService `json:"service" firestore:"service" bson:"service"`
	ExternalID string          `json:"external_id" firestore:"external_id" bson:"external_id"`
}
