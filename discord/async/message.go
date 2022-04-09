package async

type Unknown interface{}

type ContentType int

const (
	ContentTypeText ContentType = iota
	ContentTypeImage
	ContentTypeEmbed
	ContentTypeReaction
)

type Message struct {
	Data     Unknown `json:"data"`
	Locale   string  `json:"locale"`
	Metadata Unknown `json:"metadata"`
}
