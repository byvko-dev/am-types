package users

import (
	"time"

	"github.com/byvko-dev/am-types/discord/v1"
	"github.com/byvko-dev/am-types/wargaming/v1"
)

type GlobalUser struct {
	ID       string             `json:"id" firestore:"id" bson:"id"`
	Features AdditionalFeatures `json:"features" firestore:"features" bson:"features"`

	Profiles UserProfiles `json:"profiles" firestore:"profiles" bson:"profiles"`

	LastSeen  time.Time `json:"last_seen" firestore:"last_seen" bson:"last_seen"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at" bson:"updated_at"`
}

type UserProfiles struct {
	Discord   discord.DiscordProfile     `json:"discord" firestore:"discord" bson:"discord"`
	Wargaming wargaming.WargamingProfile `json:"wargaming" firestore:"wargaming" bson:"wargaming"`
}

type AdditionalFeatures struct {
	EnhancedSecurity bool `json:"enhanced_security" firestore:"enhanced_security" bson:"enhanced_security"` // Does extra checks when saving settings
	PremiumBypass    bool `json:"premium_bypass" firestore:"premium_bypass" bson:"premium_bypass"`          // Bypasses premium checks
	PreviewAccess    bool `json:"preview_access" firestore:"preview_access" bson:"preview_access"`          // Allows access to preview features
}

func (u *GlobalUser) Verify() bool {
	return u.ID != ""
}
