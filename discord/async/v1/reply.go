package async

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	api "github.com/byvko-dev/am-types/api/generic/v1"
)

type ContentType int

const (
	ContentTypeText ContentType = iota
	ContentTypeImage
	ContentTypeEmbed
	ContentTypeReaction
)

type Reply struct {
	Data  []ReplyChunk      `json:"data"`
	Error api.ResponseError `json:"error"`
}

type ReplyChunk struct {
	Data        Unknown `json:"data"`
	ContentType `json:"contentType"`
}

func (r *Reply) Encode() []byte {
	payload, _ := json.Marshal(r)
	return payload
}

func (chunk *ReplyChunk) Decode() (interface{}, int, error) {
	switch chunk.ContentType {
	case ContentTypeText:
		text, ok := chunk.Data.(string)
		if !ok {
			return nil, 0, ErrInvalidReply
		}
		return text, int(ContentTypeText), nil

	case ContentTypeImage:
		encoded, ok := chunk.Data.(string)
		if !ok {
			return nil, 0, ErrInvalidReply
		}
		raw, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return nil, 0, ErrInvalidReply
		}
		var image discordgo.File
		image.Reader = bytes.NewReader(raw)
		image.Name = "aftermath.link.png"
		image.ContentType = "image/png"
		return image, int(ContentTypeImage), nil

	case ContentTypeEmbed:
		return nil, 0, ErrResponseTypeNotImplemented

	case ContentTypeReaction:
		reaction, ok := chunk.Data.(string)
		if !ok {
			return nil, 0, ErrInvalidReply
		}
		return reaction, int(ContentTypeReaction), nil

	default:
		return nil, 0, ErrInvalidReply
	}
}
