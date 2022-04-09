package router

import (
	"github.com/bwmarrin/discordgo"
	"github.com/byvko-dev/am-types/discord/async"
)

type ReplyMode int

const (
	ReplyModeChannel ReplyMode = iota
	ReplyModeDirectMessage
	ReplyModeNone
)

type Command struct {
	Hidden      bool               `json:"hidden"`
	ReactionAck bool               `json:"reactionAck"`
	Constraits  CommandConstraints `json:"constraints"`

	ReplyMode     `json:"replyMode"`
	ArgumentsType async.MessageDataType `json:"argumentsType"`

	Name            string   `json:"name"`
	Aliases         []string `json:"aliases"`
	HelpDescription string   `json:"description"`

	MessageSettings     `json:"messageSettings"`
	InteractionSettings `json:"interaction_settings"`

	Handler func(*Command, *HandlerPayload) error `json:"-"`
}

func (c *Command) MakeInteractionSettings() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.InteractionSettings.Name,
		Description: c.InteractionSettings.DescriptionTag,
		Type:        discordgo.ChatApplicationCommand,
		Options:     c.GenerateOptions(),
	}
}

func (c *Command) MakeRemotePayload(input *HandlerPayload) (interface{}, error) {
	switch c.ArgumentsType {
	case async.Empty:
		return nil, nil

	case async.DebugRequest:
		payload := make(map[string]interface{})
		payload["command"] = c.Name
		payload["args"] = input.Arguments
		return payload, nil

	default:
		return nil, ErrInvalidCommandArgsType
	}
}

type CommandConstraints struct {
	GlobalAdminOnly            bool     `json:"globalAdminOnly"`
	UserPremissionTagsRequired []string `json:"userPremissionTagsRequired"`

	GuildOnly         bool `json:"guildOnly"`
	DirectMessageOnly bool `json:"directMessageOnly"`
}

type MessageSettings struct {
	Name string `json:"name"`

	NameTag        string `json:"nameTag"`
	DescriptionTag string `json:"descriptionTag"`

	Arguments []MessageArgument `json:"arguments"`
}

func (s *MessageSettings) Empty() bool {
	return s == &MessageSettings{}
}

type MessageArgument struct {
	Name           string `json:"name"`
	NameTag        string `json:"nameTag"`
	DescriptionTag string `json:"descriptionTag"`
	Required       bool   `json:"required"`
}

type InteractionSettings struct {
	Name string `json:"name"`

	NameTag        string `json:"nameTag"`
	DescriptionTag string `json:"descriptionTag"`

	Options []InteractionOption `json:"options"`
}

func (s *InteractionSettings) Empty() bool {
	return s == &InteractionSettings{}
}

type InteractionOption struct {
	Name string `json:"name"`

	NameTag        string `json:"nameTag"`
	DescriptionTag string `json:"descriptionTag"`

	Type     string `json:"type"`
	Required bool   `json:"required"`
}

func (s *InteractionSettings) GenerateOptions() []*discordgo.ApplicationCommandOption {
	var options []*discordgo.ApplicationCommandOption
	for _, option := range s.Options {
		options = append(options, &discordgo.ApplicationCommandOption{
			Type:         discordgo.ApplicationCommandOptionString,
			Name:         option.NameTag,
			Description:  option.DescriptionTag,
			Required:     option.Required,
			Autocomplete: true,
		})
	}
	return options
}
