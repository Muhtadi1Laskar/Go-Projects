package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

// SummarizeText returns top N important sentences
func SummarizeText(text string, numSentences int) []string {
	// Preprocess and split into sentences
	sentences := splitIntoSentences(text)
	words := tokenize(text)

	// Calculate word frequencies (excluding stop words)
	wordFreq := calculateWordFrequencies(words)

	// Score sentences based on word frequencies
	sentenceScores := scoreSentences(sentences, wordFreq)

	// Sort sentences by score and select top N
	return selectTopSentences(sentences, sentenceScores, numSentences)
}

func splitIntoSentences(text string) []string {
	// Split on .!? followed by whitespace or end of string
	re := regexp.MustCompile(`[.!?]\s+|$`)
	split := re.Split(text, -1)

	var sentences []string
	for _, s := range split {
		if trimmed := strings.TrimSpace(s); trimmed != "" {
			sentences = append(sentences, trimmed)
		}
	}
	return sentences
}

func tokenize(text string) []string {
	// Remove punctuation and convert to lowercase
	re := regexp.MustCompile(`[^\w\s]`)
	cleanText := re.ReplaceAllString(strings.ToLower(text), "")
	return strings.Fields(cleanText)
}

func calculateWordFrequencies(words []string) map[string]float64 {
	stopWords := map[string]bool{
		"the": true, "and": true, "of": true, "to": true, "a": true,
		"in": true, "that": true, "is": true, "are": true, "it": true,
	}

	freq := make(map[string]float64)
	for _, word := range words {
		if !stopWords[word] {
			freq[word]++
		}
	}

	// Normalize frequencies
	maxFreq := 0.0
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}
	}

	if maxFreq > 0 {
		for word := range freq {
			freq[word] /= maxFreq
		}
	}
	return freq
}

func scoreSentences(sentences []string, wordFreq map[string]float64) []float64 {
	scores := make([]float64, len(sentences))

	for i, sentence := range sentences {
		words := tokenize(sentence)
		totalScore := 0.0
		validWords := 0

		for _, word := range words {
			if score, exists := wordFreq[word]; exists {
				totalScore += score
				validWords++
			}
		}

		if validWords > 0 {
			// Normalize by sentence length to avoid favoring long sentences
			scores[i] = totalScore / math.Sqrt(float64(validWords))
		}
	}
	return scores
}

func selectTopSentences(sentences []string, scores []float64, n int) []string {
	type sentenceScore struct {
		sentence string
		score    float64
	}

	var ss []sentenceScore
	for i, sentence := range sentences {
		ss = append(ss, sentenceScore{sentence, scores[i]})
	}

	// Sort by descending score
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].score > ss[j].score
	})

	// Select top N sentences while preserving original order
	selected := make(map[int]bool)
	for i := 0; i < n && i < len(ss); i++ {
		selected[i] = true
	}

	var result []string
	for i, sentence := range sentences {
		if selected[i] {
			result = append(result, sentence)
		}
	}
	return result
}

func readFile(PATH string) (string, error) {
	var builder strings.Builder

	file, err := os.Open(PATH)
	if err != nil {
		return "", fmt.Errorf("failed to open the file: %v/n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error while reading file: %v", err)
	}

	return builder.String(), nil
}

func main() {
	var rootPath string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects"
	documentOne, _ := readFile(rootPath + "/Plagrism-Checker/document1.txt")

	summary := SummarizeText(documentOne, 5) // Get 2-sentence summary
	fmt.Println("Summary:")
	for _, sentence := range summary {
		fmt.Println("-", sentence)
	}
}
