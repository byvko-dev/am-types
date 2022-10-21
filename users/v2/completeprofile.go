package users

import (
	"time"
)

type CompleteProfile struct {
	ID string `json:"id" firestore:"id" bson:"_id,omitempty"`

	Locale string `json:"locale" firestore:"locale" bson:"locale"`

	Roles          Role           `json:"roles" firestore:"roles" bson:"roles"`
	Features       UserFeatures   `json:"features" firestore:"features" bson:"features"`
	Customizations Customizations `json:"customizations" firestore:"customizations" bson:"customizations"`

	ExternalProfiles []ExternalProfile `json:"external_profiles" firestore:"external_profiles" bson:"external_profiles"`

	LastSeen  time.Time `json:"last_seen" firestore:"last_seen" bson:"last_seen"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at" bson:"updated_at"`
}

func (u *CompleteProfile) Verify() bool {
	return u.ID != ""
}

func (u *CompleteProfile) ToUserCheck(banData *UserBan) UserCheck {
	services := make([]ExternalProfileID, 0, len(u.ExternalProfiles))
	for _, profile := range u.ExternalProfiles {
		services = append(services, profile.ExternalProfileID)
	}
	check := UserCheck{
		Locale:         u.Locale,
		Features:       u.Features,
		Customizations: u.Customizations,
		LinkedServices: services,
	}
	if banData != nil {
		check.Banned = time.Now().Before(banData.Expiration) && !banData.Lifted
		check.BanReason = banData.Reason
		check.BanNotified = banData.Notified
	}
	return check
}
