package main

import (
	"fmt"
	"bufio"
	"os"
	"hash/fnv"
	"strings"
)

func generateShingles(text string, n int) []string {
	shingles := []string{}
	for i := 0; i <= len(text)-n; i++ {
		shingles = append(shingles, text[i:i+n])
	}
	return shingles
}

func hashValue(shingle string, seed uint32) uint32 {
	h := fnv.New32a()
    h.Write([]byte(shingle))
    a := seed * 12345 + 67890
    b := seed * 54321 + 9876
    return (a * h.Sum32() + b) % 15485863
}

func computeMinHashSignature(shingles []string, numHashFunctions int) []uint32 {
	signature := make([]uint32, numHashFunctions)
	for i := 0; i < numHashFunctions; i++ {
		seed := uint32(i)
		minHash := uint32(^uint32(0))
		for _, shingle := range shingles {
			hash := hashValue(shingle, seed)
			if hash < minHash {
				minHash = hash
			}
		}
		signature[i] = minHash
	}
	return signature
}

func computeJaccardSimilarity(signature1, signature2 []uint32) float64 {
	matches := 0
	for i := range signature1 {
		if signature1[i] == signature2[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(signature1))
}

func readFiless(PATH string) (string, error) {
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
	documentOne, _ := readFiless(rootPath + "/Plagrism-Checker/document1.txt")
	documentTwo, _ := readFiless(rootPath + "/Plagrism-Checker/document2.txt")

	shingles1 := generateShingles(documentOne, 3)
	shingles2 := generateShingles(documentTwo, 3)

	signature1 := computeMinHashSignature(shingles1, 100)
	signature2 := computeMinHashSignature(shingles2, 100)

	similarity := computeJaccardSimilarity(signature1, signature2)
	fmt.Printf("Approximate Jaccard Similarity: %.2f%%\n", similarity*100)
}
