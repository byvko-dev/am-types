package router

import "github.com/andersfylling/disgord"

type RegisterOptions struct {
	AddSlashCommand  bool `json:"add_slash_command" bson:"add_slash_command"`
	AddPrefixCommand bool `json:"add_prefix_command" bson:"add_prefix_command"`

	Name        string                              `json:"name" bson:"name"`
	Description string                              `json:"description" bson:"description"`
	Aliases     []string                            `json:"aliases" bson:"aliases"`
	Arguments   []*disgord.ApplicationCommandOption `json:"arguments" bson:"arguments"`

	Permissions Permissions `json:"permissions" bson:"permissions"`

	Handler      func(Command) error  `json:"-" bson:"-"`
	ErrorHandler func(Command, error) `json:"-" bson:"-"`
}
