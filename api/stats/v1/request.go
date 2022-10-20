package stats

import "github.com/byvko-dev/am-types/wargaming/v2/glossary"

type RequestPayload struct {
	AccountID int    `json:"account_id" bson:"account_id"`
	Realm     string `json:"realm" bson:"realm"`
	Days      int    `json:"days" bson:"days"`

	Locale string `json:"locale" bson:"locale"`
}

func (r *RequestPayload) ParseLocale() string {
	return glossary.GetLocale(r.Locale)
}
