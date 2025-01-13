package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func generateShingle(data string, size int) []string {
	var result []string
	var formattedString string = removeStopWord(data)

	for i := 0; i < len(formattedString)-size; i++ {
		result = append(result, formattedString[i:i+size])
	}
	return result
}

func hashShingles(data []string) []string {
	var hashes []string

	for _, shingle := range data {
		hash := sha256.New().Sum([]byte(shingle))
		hashes = append(hashes, fmt.Sprintf("%x", hash))
	}
	return hashes
}

func jaccardSimilarity(set1, set2 map[string]struct{}) float64 {
	intersection := 0
	union := len(set1) + len(set2)

	for key := range set1 {
		if _, exists := set2[key]; exists {
			intersection++
		}
	}
	return float64(intersection) / float64(union-intersection)
}

func compareDocuments(doc1, doc2 string, shingleSize int) float64 {
	shinglesOne := generateShingle(doc1, shingleSize)
	shingleTwo := generateShingle(doc2, shingleSize)

	hashSetOne := make(map[string]struct{})
	hashSetTwo := make(map[string]struct{})

	for _, hash := range hashShingles(shinglesOne) {
		hashSetOne[hash] = struct{}{}
	}

	for _, hash := range hashShingles(shingleTwo) {
		hashSetTwo[hash] = struct{}{}
	}

	return jaccardSimilarity(hashSetOne, hashSetTwo)
}

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

// func main() {
// 	var rootPath string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects"
// 	documentOne, _ := readFile(rootPath + "/Plagrism-Checker/document1.txt")
// 	documentTwo, _ := readFile(rootPath + "/Plagrism-Checker/document2.txt")
// 	var size int = 3

// 	var similarity float64 = compareDocuments(documentOne, documentTwo, size)

// 	fmt.Printf("Similarity: %.2f%%\n", similarity*100)
// }
