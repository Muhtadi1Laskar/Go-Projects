package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileInformation struct {
	Path  string
	Name  string
	Size  int64
	Count int
}

func getCodeInfo(root string, extension string) ([]FileInformation, int) {
	var result []FileInformation
	var totalLines int = 0

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(path) == extension {
			lineCount := countCodeLines(path)
			result = append(result, FileInformation{
				Path:  path,
				Name:  info.Name(),
				Size:  info.Size(),
				Count: lineCount,
			})

			totalLines += lineCount
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return result, totalLines
}

func countCodeLines(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	var count int = 0
	for scanner.Scan() {
		count += 1
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("Failed to scan the file: %v", err)
	}

	return count
}

func main() {
	var filePathString string = "C:/Users/laska/OneDrive/Documents/Coding/JavaScript/Node-Projects"

	files, totalLines := getCodeInfo(filePathString, ".js")

	for _, value := range files {
		fmt.Println("Path: ", value.Path)
		fmt.Println("Name: ", value.Name)
		fmt.Println("Size: ", value.Size)
		fmt.Println("Count: ", value.Count)
		fmt.Println()
	}

	fmt.Println("Total Lines of Code: ", totalLines)
}
