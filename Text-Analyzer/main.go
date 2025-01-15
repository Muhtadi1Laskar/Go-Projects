package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var stopWordsSet = map[string]struct{}{
	"a": {}, "an": {}, "this": {}, "the": {}, "is": {}, "are": {}, "was": {}, "were": {}, "will": {}, "be": {},
	"in": {}, "on": {}, "at": {}, "of": {}, "for": {}, "to": {}, "from": {}, "with": {},
	"and": {}, "or": {}, "but": {}, "not": {}, "if": {}, "then": {}, "else": {},
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {}, "we": {}, "they": {}, "my": {}, "your": {}, "his": {}, "her": {}, "its": {}, "our": {}, "their": {},
	"couldve": {}, "couldnt": {}, "wouldnt": {}, "shouldnt": {}, "wasnt": {}, "wont": {}, "shallnt": {}, "didnt": {}, "weev": {}, "im": {},
}

var BASE_PATH string = "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Text-Analyzer/"

func countWords(text string) int {
	text = strings.TrimSpace(text)
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

func countLetters(text string) int {
	count := 0

	for _, char := range text {
		if unicode.IsLetter(char) && !isWhiteSpace(char) {
			count++
		}
	}
	return count
}

func countSentences(text string) int {
	count := 0
	insideSentence := false

	for _, char := range text {
		if char == '.' || char == '?' || char == '!' {
			if insideSentence {
				count++
				insideSentence = false
			}
		} else if char == ' ' || char == '\n' || char == '\t' {
			insideSentence = true
		}
	}

	if insideSentence {
		count++
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
	case "letter-count":
		return countLetters(text)
	case "sentence-count":
		return countSentences(text)
	default:
		return -1
	}
}

func removePunctuation(text string) string {
	var builder strings.Builder

	for _, word := range text {
		if word == '-' || word == '\'' || !unicode.IsPunct(word) {
			builder.WriteRune(word)
		}
	}
	return builder.String()
}

func calculateFreq(text string) map[string]int {
	words := cleanText(text)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func tokenize(text string) []string {
	var tokens []string
	var wordBuilder strings.Builder

	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			wordBuilder.WriteRune(unicode.ToLower(char))
		} else if wordBuilder.Len() > 0 {
			tokens = append(tokens, wordBuilder.String())
			wordBuilder.Reset()
		}
	}

	if wordBuilder.Len() > 0 {
		tokens = append(tokens, wordBuilder.String())
	}

	return tokens
}

func removeStopWords(text []string) []string {
	var filteredStr []string

	for _, word := range text {
		if _, exists := stopWordsSet[word]; !exists {
			filteredStr = append(filteredStr, word)
		}
	}
	return filteredStr
}

func stemWord(text string) string {
	suffixes := []string{"ly", "ing", "ed"}

	for _, suffix := range suffixes {
		if strings.HasSuffix(text, suffix) {
			text = strings.TrimSuffix(text, suffix)
			break
		}
	}
	return text
}

func sufixStripping(text string) string {
	words := strings.Fields(text)

	for key, word := range words {
		words[key] = stemWord(word)
	}
	return strings.Join(words, " ")
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

func averageWordLenght(data string) float64 {
	return float64(count(data, "letter-count") / count(data, "word-count"))
}

func loadWords(filePath string) map[string]bool {
	data, _ := readFile(filePath)
	wordMap := make(map[string]bool)

	words := strings.Split(strings.TrimSpace(data), ",")

	for _, word := range words {
		wordMap[strings.ToLower(strings.TrimSpace(word))] = true
	}
	return wordMap
}

func sentimentAnalysis(text string) string {
	positiveMap := loadWords(BASE_PATH + "Data/positiveWords.txt")
	negativeMap := loadWords(BASE_PATH + "Data/negativeWords.txt")
	tokens := tokenize(text)
	positiveCount, negativeCount := 0, 0

	for _, token := range tokens {
		if positiveMap[token] {
			positiveCount++
		} else if negativeMap[token] {
			negativeCount++
		}
	}

	fmt.Println(positiveCount, negativeCount)

	if positiveCount > negativeCount {
		return "The given text is positive"
	} else if positiveCount < negativeCount {
		return "The given text is negative"
	} else {
		return "The text is neutral"
	}
}

func cleanText(text string) []string {
	formattedStr := removePunctuation(text)
	allWords := tokenize(formattedStr)
	words := removeStopWords(allWords)

	return words
}

func main() {
	var path string = BASE_PATH + "Data/text.txt"
	data, _ := readFile(path)
	freq := calculateFreq(data)

	fmt.Println("Word Count:", count(data, "word-count"))
	fmt.Println("Character Count:", count(data, "character-count"))
	fmt.Println("Letter Count: ", count(data, "letter-count"))
	fmt.Println("Sentence Count:", count(data, "sentence-count"))
	fmt.Printf("Average Word Length: %.2f%%\n", averageWordLenght(data))

	fmt.Println(sufixStripping("The dogs are running quickly and happily after playing with gaining"))

	for key, value := range freq {
		fmt.Println(key, value)
	}

	fmt.Println(sentimentAnalysis(data))
	fmt.Println(tokenize(data))
}
