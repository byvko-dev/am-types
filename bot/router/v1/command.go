package router

import (
	"github.com/andersfylling/disgord"
	"github.com/byvko-dev/am-types/users/v1"
)

type Command struct {
	Name      string   `json:"name" bson:"name"`
	Arguments []string `json:"arguments" bson:"arguments"`

	User users.CompleteProfile `json:"user" bson:"user"`

	Unsafe InternalOptions `json:"-" bson:"-"`
}

type InternalOptions struct {
	UserID  string                                        `json:"-" bson:"-"`
	Options RegisterOptions                               `json:"-" bson:"-"`
	Session disgord.Session                               `json:"-" bson:"-"`
	Message *disgord.Message                              `json:"-" bson:"-"`
	Reply   func(c Command, content ...interface{}) error `json:"-" bson:"-"`
}
