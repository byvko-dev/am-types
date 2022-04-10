package async

import "github.com/byvko-dev/am-types/users/v2"

type Unknown interface{}

type CommandType int

const (
	CommandTypeMessage CommandType = iota
	CommandTypeInteraction
	CommandTypeReaction
)

type Message struct {
	Type CommandType `json:"commandType"`

	Files     []string `json:"files"`
	Prefix    string   `json:"prefix"`
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`

	Locale   string                `json:"locale"`
	Metadata Unknown               `json:"metadata"`
	User     users.CompleteProfile `json:"user"`
}
