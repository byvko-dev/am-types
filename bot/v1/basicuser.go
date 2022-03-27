package bot

type BasicUser struct {
	ID     string `json:"id" bson:"id"`
	Locale string `json:"locale" bson:"locale"`
	UserBan

	FeaturesEnabled    []string `json:"featuresEnabled" bson:"featuresEnabled"`
	FeaturesRestricted []string `json:"featuresRestricted" bson:"featuresRestricted"`
}
