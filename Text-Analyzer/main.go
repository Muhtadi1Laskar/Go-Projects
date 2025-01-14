package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countWords(text string) int {
	return len(strings.Fields(text))
}

func countCharacters(text string) int {
	count := 0

	for _, char := range text {
		if !isWhiteSpace(char) {
			count++
		}
	}
	return count
}

func countSentences(text string) int {
	count := 0
	sentenceDelimeters := map[rune]bool{
		'.': true,
		'!': true,
		'?': true,
	}
	insideSentence := false

	for _, char := range text {
		if sentenceDelimeters[char] {
			if insideSentence {
				count++
				insideSentence = false
			}
		} else if !isWhiteSpace(char) {
			insideSentence = true
		}
	}
	return count
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t'
}

func count(text, operation string) int {
	text = strings.TrimSpace(text)
	operation = strings.ToLower(operation)

	switch operation {
	case "word-count":
		return countWords(text)
	case "character-count":
		return countCharacters(text)
	case "sentence-count":
		return countSentences(text)
	default:
		return -1
	}
}

func removePunctuation(text string) string {
	var builder strings.Builder

	for _, word := range text {
		if word == '-' || word == '\'' ||!unicode.IsPunct(word) {
			builder.WriteRune(word)
		}
	}
	return builder.String()
}

func calculateFreq(text string) map[string]int {
	formattedStr := removePunctuation(text)
	words := tokenize(formattedStr)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func tokenize(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
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
	var path string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Text-Analyzer/text.txt"
	data, _ := readFile(path)
	freq := calculateFreq(data)

	fmt.Println("Word Count:", count(data, "word-count"))
	fmt.Println("Character Count:", count(data, "character-count"))
	fmt.Println("Sentence Count:", count(data, "sentence-count"))

	for key, value := range freq {
		fmt.Println(key, value)
	}
}
