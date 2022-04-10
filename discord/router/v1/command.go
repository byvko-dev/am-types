package router

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	str "github.com/byvko-dev/am-core/helpers/strings"
	"github.com/byvko-dev/am-types/discord/async/v1"
)

type ReplyMode string

const (
	ReplyModeChannel       ReplyMode = "channel"
	ReplyModeDirectMessage ReplyMode = "directMessage"
	ReplyModeReaction      ReplyMode = "reaction"
	ReplyModeNone          ReplyMode = "none"
)

type Command struct {
	Hidden      bool               `json:"hidden"`
	ReactionAck bool               `json:"reactionAck"`
	Constraits  CommandConstraints `json:"constraints"`

	ReplyMode `json:"replyMode" validate:"required"`
	TopicName string `json:"topicName" validate:"required"`

	HelpDescription string `json:"description"`

	MessageSettings     `json:"messageSettings"`
	InteractionSettings `json:"interaction_settings"`

	Payload      *HandlerPayload       `json:"payload"`
	Handler      func(*Command) error  `json:"-"`
	ErrorHandler func(*Command, error) `json:"-"`
}

func (c *Command) MakeInteractionSettings() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        c.InteractionSettings.Name,
		Description: c.InteractionSettings.DescriptionTag,
		Type:        discordgo.ChatApplicationCommand,
		Options:     c.GenerateOptions(),
	}
}

func (c *Command) MakeRemotePayload(input *HandlerPayload) ([]byte, error) {
	var request async.Message
	request.Files = input.Files
	request.Metadata = input.Metadata
	request.Arguments = input.Arguments
	request.Locale = str.Or(input.User.Locale, input.Metadata.GuildLocale)
	return json.Marshal(request)
}

type CommandConstraints struct {
	GlobalAdminOnly            bool     `json:"globalAdminOnly"`
	UserPremissionTagsRequired []string `json:"userPremissionTagsRequired"`

	GuildOnly         bool `json:"guildOnly"`
	DirectMessageOnly bool `json:"directMessageOnly"`
}

type MessageSettings struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`

	NameTag        string `json:"nameTag"`
	DescriptionTag string `json:"descriptionTag"`

	Arguments []MessageArgument `json:"arguments"`
}

func (s *MessageSettings) Empty() bool {
	return s == &MessageSettings{}
}

type MessageArgument struct {
	Name    string   `json:"name"`
	Aliases []string `json:"aliases"`

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
