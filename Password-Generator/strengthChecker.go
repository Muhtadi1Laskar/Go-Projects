package main

type CommonPatterns struct {
	Sequentials string `json:"sequentials"`
	Repeated    string `json:"repeated"`
}

type PasswordAnalysis struct {
	Length         int            `json:"length"`
	HasLowerCase   bool           `json:"hasLowerCase"`
	HasUpperCase   bool           `json:"hasUpperCase"`
	HasNumber      bool           `json:"hasNumber"`
	HasSymbols     bool           `json:"hasSymbols"`
	CommonPatterns CommonPatterns `json:"commonPatterns"`
	Strength       string         `json:"strength"`
	Score          float64        `json:"score"`
}
