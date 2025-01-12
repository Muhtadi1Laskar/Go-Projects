package main

import (
	"crypto/sha256"
	"fmt"
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

func main() {
	var documentOne string = "This is a fucking string"
	var size int = 3

	shingles := generateShingle(documentOne, size)
	hashes := hashShingles(shingles)

	fmt.Println(hashes)
}