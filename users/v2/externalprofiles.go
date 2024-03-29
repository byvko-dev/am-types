package users

type ExternalService string
type ExternalServiceConfig struct {
	Name   ExternalService `json:"name" firestore:"name" bson:"name"`
	Unique bool            `json:"unique" firestore:"unique" bson:"unique"`
}

var (
	ExternalServiceWargaming = ExternalServiceConfig{
		Name:   "wargaming",
		Unique: false,
	}
	ExternalServiceDiscord = ExternalServiceConfig{
		Name:   "discord",
		Unique: true,
	}
	// Indexes need to match
	ValidExternalServiceNames = []ExternalService{ExternalServiceWargaming.Name, ExternalServiceDiscord.Name}
	ValidExternalServices     = []ExternalServiceConfig{ExternalServiceWargaming, ExternalServiceDiscord}
)

type ExternalConnection struct {
	ID                string `bson:"_id,omitempty" json:"id"`
	ExternalProfileID `bson:",inline" json:",inline"`

	UserID       string            `json:"user_id" bson:"user_id"`
	AccessToken  string            `json:"access_token" firestore:"access_token" bson:"access_token"`
	RefreshToken string            `json:"refresh_token" firestore:"refresh_token" bson:"refresh_token"`
	Meta         map[string]string `json:"meta" firestore:"meta" bson:"meta"`
}

type ExternalProfileID struct {
	Service    ExternalService `json:"service" firestore:"service" bson:"service"`
	ExternalID string          `json:"external_id" firestore:"external_id" bson:"external_id"`
	Verified   bool            `json:"verified" firestore:"-" bson:"-"`
}
