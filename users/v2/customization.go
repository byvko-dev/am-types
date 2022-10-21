package users

type Customizations struct {
	Pins             []Pin  `json:"pins" bson:"pins"`
	BackgroundImage  string `json:"backgroundImage" bson:"backgroundImage"`
	CustomSettingsID string `json:"customSettingsId" bson:"customSettingsId"`
}
