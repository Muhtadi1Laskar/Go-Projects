package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

type CommonPatterns struct {
	Sequential bool `json:"sequentials"`
	Repeated   bool `json:"repeated"`
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

func CheckStrength(password string) PasswordAnalysis {
	var score float64

	analysis := PasswordAnalysis{
		Length:       len(password),
		HasLowerCase: regexp.MustCompile(`[a-z]`).MatchString(password),
		HasUpperCase: regexp.MustCompile(`[A-Z]`).MatchString(password),
		HasNumber:    regexp.MustCompile(`\d`).MatchString(password),
		HasSymbols:   regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>?]`).MatchString(password),
		CommonPatterns: CommonPatterns{
			Sequential: regexp.MustCompile(`(abc|bcd|cde|def|efg|fgh|ghi|hij|ijk|jkl|klm|lmn|mno|nop|opq|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz|012|123|234|345|456|567|678|789)`).MatchString(password),
			Repeated:   regexp.MustCompile(`(.)\\1{2,}`).MatchString(password),
		},
	}

	if analysis.Length >= 8 {
		score += 1
	}
	if analysis.Length >= 12 {
		score += 1
	}
	if analysis.Length >= 16 {
		score += 1
	}

	if analysis.HasLowerCase {
		score += 1
	}
	if analysis.HasUpperCase {
		score += 1
	}
	if analysis.HasNumber {
		score += 1
	}
	if analysis.HasSymbols {
		score += 1
	}

	if analysis.CommonPatterns.Sequential {
		score -= 1
	}
	if analysis.CommonPatterns.Repeated {
		score -= 1
	}

	switch {
	case score <= 3:
		analysis.Strength = "weak"
	case score <= 5:
		analysis.Strength = "medium"
	case score <= 7:
		analysis.Strength = "strong"
	default:
		analysis.Strength = "very strong"
	}

	if score < 0 {
		score = 0
	}
	if score > 10 {
		score = 10
	}

	analysis.Score = score * 1.4
	if analysis.Score > 10 {
		analysis.Score = 10
	}

	return analysis
}

func main() {
	password := "Hello890!@#World"
	strengthStat := CheckStrength(password)

	bytes, err := json.MarshalIndent(strengthStat, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal struct to JSON: %v", err)
	}
	fmt.Println(string(bytes))
}