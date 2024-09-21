package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func readFile() string {
	var message string;
	file, err := os.Open("files/one-text.txt");
	if err != nil {
		log.Fatal(err);
	}
	defer file.Close();

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		message += scanner.Text() + "\n";
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err);
	}
	return message;
}

func main() {
	var message string = readFile();

	fmt.Println(message);
}