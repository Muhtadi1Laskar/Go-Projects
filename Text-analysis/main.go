package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var PATH string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Text-analysis/dummy.txt"

func readFile(PATH string) (string, error) {
	var builder strings.Builder

	file, err := os.Open(PATH)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
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

func capitalize(data string) string {
	return strings.ToUpper(data)
}

func deCapitalize(data string) string {
	return strings.ToLower(data)
}

func capitalizeFirst(data string) string {
	var result strings.Builder
	words := strings.Fields(data)

	for _, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			result.WriteRune(unicode.ToUpper(runes[0]))
			result.WriteString(strings.ToLower(string(runes[1:])))
		}
		result.WriteRune(' ')
	}
	return strings.TrimSpace(result.String())
}

func frequency(data string) map[string]int {
	hashMap := make(map[string]int)
	format := strings.ReplaceAll(data, " ", "")
	data = strings.ToLower(format)

	for _, char := range data {
		hashMap[string(char)]++
	}
	return hashMap
}


func main() {
	text, err := readFile(PATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(capitalizeFirst(text))
	fmt.Println(frequency(text))
}