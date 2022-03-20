package api

type ResponseWithError struct {
	Data  interface{}   `json:"data" firestore:"data" bson:"data"`
	Error ResponseError `json:"error" firestore:"error" bson:"error"`
}

type ResponseError struct {
	Error   string `json:"error" firestore:"error" bson:"error"`
	Context string `json:"context" firestore:"context" bson:"context"`
}
