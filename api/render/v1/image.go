package render

type ImageWithMetrics struct {
	Metrics struct {
		ConsoleLogs  []ConsoleLog `json:"consoleLogs" bson:"consoleLogs"`
		TimeToFinish float64      `json:"timeToFinish" bson:"timeToFinish"`
		TimeToWarmup float64      `json:"timeToWarmup" bson:"timeToWarmup"`
		TimeToCache  float64      `json:"timeToCache" bson:"timeToCache"`
	} `json:"metrics" bson:"metrics"`
	Image []byte `json:"image" bson:"image"`
}

type ConsoleLog struct {
	Level string   `json:"level" bson:"level"`
	Lines []string `json:"lines" bson:"lines"`
}
