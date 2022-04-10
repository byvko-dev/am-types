package router

import (
	"github.com/bwmarrin/discordgo"
	"github.com/byvko-dev/am-types/users/v2"
)

type InteractionType string

const (
	Message            InteractionType = "message"
	Reaction           InteractionType = "reaction"
	FileUpload         InteractionType = "fileUpload"
	ApplicationCommand InteractionType = "applicationCommand"
)

type HandlerPayload struct {
	session *discordgo.Session `json:"-"`

	Metadata        `json:"metadata"`
	InteractionType `json:"interactionType"`
	Name            string                              `json:"name"`
	Prefix          string                              `json:"prefix"`
	Arguments       []string                            `json:"arguments"`
	Files           []string                            `json:"files"`
	User            users.CompleteProfile               `json:"user"`
	Reply           func(...interface{}) error          `json:"-"`
	Printer         func(string, ...interface{}) string `json:"-"`

	errorHandler func(*HandlerPayload, error) `json:"-"`
}

func (p *HandlerPayload) HandleError(err error) {
	p.errorHandler(p, err)
}

func (p *HandlerPayload) SetErrorHandler(handler func(*HandlerPayload, error)) {
	p.errorHandler = handler
}

func (p *HandlerPayload) SetSession(s *discordgo.Session) {
	p.session = s
}

func (p *HandlerPayload) UnsafeSession() *discordgo.Session {
	return p.session
}

func (p *HandlerPayload) SetGuildLocale(s *discordgo.Session) {
	if p.GuildID == "" || p.GuildID == "0" {
		return
	}
	guild, err := s.State.Guild(p.GuildID)
	if err != nil || guild == nil {
		p.GuildLocale = "en"
		return
	}
	p.GuildLocale = p.guildLocaleToLocale(guild.PreferredLocale)
}

func (p *HandlerPayload) guildLocaleToLocale(guildLocale string) string {
	switch guildLocale {
	case "ru-RU":
		return "ru"
	default:
		return "en"
	}
}
