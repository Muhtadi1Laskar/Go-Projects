package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"crypto/sha512"
	"strings"
)

func readFile() string {
	var builder strings.Builder;
	file, err := os.Open("files/one-text.txt");
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		builder.WriteString(scanner.Text() + "\n");
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err);
	}
	return builder.String();
}

func hashFunction(text string) []byte {
	byteMessage := []byte(text);
	hash := sha512.New();
	hash.Write(byteMessage);

	return hash.Sum(nil);
}

func main() {
	var message string = readFile();
	hashedMsg := hashFunction(message);

    fmt.Printf("SHA256: %x\n", hashedMsg);
}