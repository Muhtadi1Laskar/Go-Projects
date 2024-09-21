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

func main() {
	var message string = readFile();
	byteMessage := []byte(message);
	hash := sha512.New();
	hash.Write(byteMessage)

    fmt.Printf("SHA256: %x\n", hash.Sum(nil))
}