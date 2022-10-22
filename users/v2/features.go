package users

import (
	"github.com/byvko-dev/am-core/helpers/slices"
)

type Feature string
type Features []Feature

type UserFeatures struct {
	Enabled  Features `json:"enabled" bson:"enabled"`
	Disabled Features `json:"disabled" bson:"disabled"`
}

func (u *UserFeatures) SetDefaults() {
	u.Enabled = []Feature{CustomizeSettings, CustomizeBackgrounds}
	u.Disabled = []Feature{}
}

const PreviewAccess Feature = "preview_access"
const EnhancedSecurity Feature = "enhanced_security"

const PremiumBypass Feature = "premium_bypass"
const CustomizeSettings Feature = "customize_settings"
const CustomizeBackgrounds Feature = "customize_backgrounds"

func (all *Features) Includes(f Feature) bool {
	return slices.Contains(*all, f) > -1
}
