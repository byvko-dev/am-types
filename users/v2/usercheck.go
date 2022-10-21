package users

type UserCheck struct {
	Banned      bool   `json:"banned" bson:"banned"`
	BanReason   string `json:"ban_reason" bson:"ban_reason"`
	BanNotified bool   `json:"ban_notified" bson:"ban_notified"`

	Roles          Roles          `json:"roles" firestore:"roles" bson:"roles"`
	Locale         string         `json:"locale" firestore:"locale" bson:"locale"`
	Features       UserFeatures   `json:"features" firestore:"features" bson:"features"`
	Customizations Customizations `json:"customizations" firestore:"customizations" bson:"customizations"`

	LinkedServices []ExternalProfileID `json:"linked_services" firestore:"linked_services" bson:"linked_services"`
}

func (u *UserCheck) GetExternalID(service ExternalService) (string, bool) {
	for _, s := range u.LinkedServices {
		if s.Service == service {
			return s.ExternalID, true
		}
	}
	return "", false
}

func (u *UserCheck) FeatureEnabled(feature string) bool {
	for _, f := range u.Features.Disabled {
		if string(f) == feature {
			return false
		}
	}
	for _, f := range u.Features.Enabled {
		if string(f) == feature {
			return true
		}
	}
	return false
}
