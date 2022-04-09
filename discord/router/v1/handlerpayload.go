package router

import (
	"github.com/bwmarrin/discordgo"
	"github.com/byvko-dev/am-types/users/v2"
)

type HandlerPayload struct {
	session *discordgo.Session
	Command *Command

	Metadata
	Name      string
	Arguments []string
	User      users.CompleteProfile
	Reply     func(...interface{}) error
	Printer   func(string, ...interface{}) string

	ErrorHandler func(*HandlerPayload, error)
}

func (p *HandlerPayload) SetSession(s *discordgo.Session) {
	p.session = s
}
