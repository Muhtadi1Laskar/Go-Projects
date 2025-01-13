package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func tokenize(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func calculateTF(words []string) map[string]float64 {
	tf := make(map[string]float64)
	totalWords := float64(len(words))

	for _, word := range words {
		tf[word]++
	}

	for word := range tf {
		tf[word] /= totalWords
	}

	return tf
}

func calculateIDF(corpus [][]string) map[string]float64 {
	idf := make(map[string]float64)
	totalDocs := float64(len(corpus))

	for _, document := range corpus {
		wordSet := make(map[string]bool)

		for _, word := range document {
			wordSet[word] = true
		}

		for word := range wordSet {
			idf[word]++
		}
	}

	for word := range idf {
		idf[word] = math.Log(totalDocs / (1 + idf[word]))
	}

	return idf
}

func calculateTFIDF(words []string, idf map[string]float64) map[string]float64 {
	tf := calculateTF(words)
	tfidf := make(map[string]float64)

	for words, tfValue := range tf {
		tfidf[words] = tfValue * idf[words]
	}

	return tfidf
}

func cosineSimilarity(vec1, vec2 map[string]float64) float64 {
	var dotProduct, magnitudeA, magnitudeB float64

	for word, value1 := range vec1 {
		value2 := vec2[word]
		dotProduct += value1 * value2
		magnitudeA += value1 * value1
	}

	for _, value := range vec2 {
		magnitudeB += value * value
	}

	magnitudeA = math.Sqrt(magnitudeA)
	magnitudeB = math.Sqrt(magnitudeB)

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return dotProduct / (magnitudeA * magnitudeB)
}

func readFiles(PATH string) (string, error) {
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

var stopWordsSets = map[string]struct{}{
	"a": {}, "an": {}, "this": {}, "the": {}, "is": {}, "are": {}, "was": {}, "were": {}, "will": {}, "be": {},
	"in": {}, "on": {}, "at": {}, "of": {}, "for": {}, "to": {}, "from": {}, "with": {},
	"and": {}, "or": {}, "but": {}, "not": {}, "if": {}, "then": {}, "else": {},
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {}, "we": {}, "they": {}, "my": {}, "your": {}, "his": {}, "her": {}, "its": {}, "our": {}, "their": {},
}

func removeStopWords(text string) string {
	formattedText := strings.Fields(strings.ToLower(text))
	filteredWords := make([]string, 0, len(formattedText))

	for _, word := range formattedText {
		if _, exists := stopWordsSets[word]; !exists {
			filteredWords = append(filteredWords, word)
		}
	}
	return strings.Join(filteredWords, " ")

}

// func main() {
// 	var rootPath string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects"
// 	dataOne, _ := readFiles(rootPath + "/Plagrism-Checker/document1.txt")
// 	dataTwo, _ := readFiles(rootPath + "/Plagrism-Checker/document2.txt")
// 	data1 := removeStopWords(dataOne)
// 	data2 := removeStopWords(dataTwo)

// 	corpus := [][]string{
// 		tokenize(data1),
// 	}

// 	inputDoc := tokenize(data2)

// 	idf := calculateIDF(corpus)

// 	inputTFIDF := calculateTFIDF(inputDoc, idf)

// 	for i, document := range corpus {
// 		corpusTFIDF := calculateTFIDF(document, idf)
// 		similarity := cosineSimilarity(inputTFIDF, corpusTFIDF)
// 		fmt.Printf("Similarity with Document %d: %.2f\n", i+1, similarity*100)
// 	}
// }