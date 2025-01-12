package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func generateShingle(data string, size int) []string {
	var result []string

	for i := 0; i < len(data)-size; i++ {
		result = append(result, data[i:i+size])
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
	"a": {}, "an": {}, "the": {}, "is": {}, "are": {}, "was": {}, "were": {}, "will": {}, "be": {},
	"in": {}, "on": {}, "at": {}, "of": {}, "for": {}, "to": {}, "from": {}, "with": {},
	"and": {}, "or": {}, "but": {}, "not": {}, "if": {}, "then": {}, "else": {},
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {}, "we": {}, "they": {}, "my": {}, "your": {}, "his": {}, "her": {}, "its": {}, "our": {}, "their": {},
}

func removeStopWords(text string) string {
	var stopWords []string = []string{"a", "an", "the", "this", "is", "are", "was", "were", "will", "be",
	"in", "on", "at", "of", "for", "to", "from", "with",
	"and", "or", "but", "not", "if", "then", "else",
	"i", "you", "he", "she", "it", "we", "they", "my", "your", "his", "her", "its", "our", "their"}

	var words []string = strings.Fields(strings.ToLower(text))
	var filteredWords []string

	for _, word := range words {
		if !contains(word, stopWords) {
			filteredWords = append(filteredWords, word)
		}
	}
	return strings.Join(filteredWords, " ")
}

func contains(word string, stopWords []string) bool {
	for _, data := range stopWords {
		if data == word {
			return true
		}
	}
	return false
}

func main() {
	var documentOne string = "This is an example sentence to remove stop words."
	var documentTwo string = "This is also a fucking string"
	var size int = 3

	fmt.Println(removeStopWords(documentOne))

	var similarity float64 = compareDocuments(documentOne, documentTwo, size)

	fmt.Printf("Similarity: %.2f%%\n", similarity * 100)
}