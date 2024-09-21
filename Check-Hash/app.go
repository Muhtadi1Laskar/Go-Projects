package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath string = "hash_output.json";

type Hash struct {
	H string `json: "hash"`
}

func saveToJSON(data Hash) error {
	file, err := os.Create(filePath);
	if err != nil {
		return err;
	}
	defer file.Close();

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err;
	}

	_, err = file.Write(jsonData);
	if err != nil {
		return err;
	}

	return nil;
}

func readJsonFile() (Hash, error) {
	var data Hash;

	file, err := os.Open(filePath);
	if err != nil {
		return data, err;
	}
	defer file.Close();

	fileContent, err := io.ReadAll(file);
	if err != nil {
		return data, err;
	}

	err = json.Unmarshal(fileContent, &data);
	if err != nil {
		return data, err;
	}

	return data, nil;
}

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

func hashFunction(text string) Hash {
	byteMessage := []byte(text);
	hash := sha512.New();
	hash.Write(byteMessage);

	hashedBytes := hash.Sum(nil);
	encodedStr := hex.EncodeToString(hashedBytes);

	return Hash{
		H: encodedStr,
	}
}

func detectChange() (bool, error) {
	message := readFile();
	hashedMessage := hashFunction(message);
	previousHash, err := readJsonFile();
	if err != nil {
		return false, err;
	}
	return previousHash.H == hashedMessage.H, nil;
}

func main() {
	// var message string = readFile();
	// hashedMsg := hashFunction(message);
	// err := saveToJSON(hashedMsg);
	// if err != nil {
	// 	log.Fatal(err);
	// }
	isChanged, err := detectChange();
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}

	fmt.Println(isChanged);

    // fmt.Printf("SHA256: %s\n", hashedMsg);
}