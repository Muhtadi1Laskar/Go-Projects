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