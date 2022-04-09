package router

type Metadata struct {
	GuildID     string `json:"guildId"`
	ChannelID   string `json:"channelId"`
	MessageID   string `json:"messageId"`
	UserID      string `json:"userId"`
	GuildLocale string `json:"guildLocale"`
}
