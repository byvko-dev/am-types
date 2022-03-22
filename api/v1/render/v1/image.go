package render

type ImageWithMetrics struct {
	Metrics struct {
		ConsoleLogs  []ConsoleLog `json:"consoleLogs"`
		TimeToFinish float64      `json:"timeToFinish"`
		TimeToWarmup float64      `json:"timeToWarmup"`
		TimeToCache  float64      `json:"timeToCache"`
	} `json:"metrics"`
	Image []byte `json:"image"`
}

type ConsoleLog struct {
	Level string   `json:"level"`
	Lines []string `json:"lines"`
}
