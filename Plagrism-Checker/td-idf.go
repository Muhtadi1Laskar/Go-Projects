package main

import (
	"fmt"
	"math"
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