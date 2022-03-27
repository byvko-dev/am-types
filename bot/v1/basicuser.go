package bot

type BasicUser struct {
	ID     string `json:"id"`
	Locale string `json:"locale"`
	UserBan

	FeaturesEnabled    []string `json:"featuresEnabled"`
	FeaturesRestricted []string `json:"featuresRestricted"`
}
