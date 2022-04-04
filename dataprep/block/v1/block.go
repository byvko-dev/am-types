package block

type BlockWithPosition struct {
	Block
	Position int `json:"position" bson:"position"`
}

type Block struct {
	Style Style `json:"style" bson:"style"`

	ContentType ContentType `json:"contentType" bson:"contentType"`
	Content     interface{} `json:"content" bson:"content"`

	Tags []string `json:"tags" bson:"tags"`
}
