package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

var stopWordsSet = map[string]struct{}{
	"a": {}, "an": {}, "this": {}, "the": {}, "is": {}, "are": {}, "was": {}, "were": {}, "will": {}, "be": {},
	"in": {}, "on": {}, "at": {}, "of": {}, "for": {}, "to": {}, "from": {}, "with": {},
	"and": {}, "or": {}, "but": {}, "not": {}, "if": {}, "then": {}, "else": {},
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {}, "we": {}, "they": {}, "my": {}, "your": {}, "his": {}, "her": {}, "its": {}, "our": {}, "their": {},
}

func removeStopWord(text string) string {
	formattedText := strings.Fields(strings.ToLower(text))
	filteredWords := make([]string, 0, len(formattedText))

	for _, word := range formattedText {
		if _, exists := stopWordsSet[word]; !exists {
			filteredWords = append(filteredWords, word)
		}
	}
	return strings.Join(filteredWords, " ")

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

// cosineSimilarity calculates the cosine similarity between two vectors
func cosineSimilaritys(vector1, vector2 []float64) float64 {
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	for i := 0; i < len(vector1); i++ {
		dotProduct += vector1[i] * vector2[i]
		magnitude1 += vector1[i] * vector1[i]
		magnitude2 += vector2[i] * vector2[i]
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0.0
	}

	return dotProduct / (magnitude1 * magnitude2)
}

// preprocessText preprocesses the text (lowercase, remove punctuation, tokenize)
func preprocessText(text string) []string {
	// Convert to lowercase
	text = strings.ToLower(text)
	text = removeStopWord(text)

	// Remove punctuation
	reg := regexp.MustCompile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	// Tokenize into words
	return strings.Fields(text)
}

// createFrequencyMap creates a frequency map of words
func createFrequencyMap(words []string) map[string]int {
	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}
	return freqMap
}

// getAllWords gets all unique words from both frequency maps
func getAllWords(freqMap1, freqMap2 map[string]int) []string {
	allWords := make(map[string]bool)
	for word := range freqMap1 {
		allWords[word] = true
	}
	for word := range freqMap2 {
		allWords[word] = true
	}

	// Convert map keys to a slice
	uniqueWords := make([]string, 0, len(allWords))
	for word := range allWords {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}

// createVector creates a vector for a given frequency map
func createVector(allWords []string, freqMap map[string]int) []float64 {
	vector := make([]float64, len(allWords))
	for i, word := range allWords {
		vector[i] = float64(freqMap[word])
	}
	return vector
}

func main() {
	var rootPath string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects"
	documentOne, _ := readFile(rootPath + "/Plagrism-Checker/document1.txt")
	documentTwo, _ := readFile(rootPath + "/Plagrism-Checker/document2.txt")

	// Preprocess the texts
	words1 := preprocessText(documentOne)
	words2 := preprocessText(documentTwo)

	// Create word frequency maps
	freqMap1 := createFrequencyMap(words1)
	freqMap2 := createFrequencyMap(words2)

	// Get all unique words
	allWords := getAllWords(freqMap1, freqMap2)

	// Create vectors
	vector1 := createVector(allWords, freqMap1)
	vector2 := createVector(allWords, freqMap2)

	// Calculate cosine similarity
	similarity := cosineSimilaritys(vector1, vector2)
	fmt.Printf("Similarity: %.2f\n", similarity)
}
