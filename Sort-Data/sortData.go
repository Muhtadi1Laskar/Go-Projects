package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

// Read CSV data from a file.
func readCSV() ([][]string, error) {
	path, err := getFilePath("output")
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Write CSV data to a file.
func writeCSV(data [][]string, filename string) {
	path, err := getFilePath(filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("cannot create the CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		log.Fatalf("cannot write to the CSV file: %v", err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalf("error flushing CSV writer: %v", err)
	}
}

// Build an integer array from CSV data.
func buildArray(data [][]string) []int {
	array := make([]int, 0, len(data))
	for _, elem := range data {
		num, _ := strconv.Atoi(elem[0])
		array = append(array, num)
	}
	return array
}

// Convert integer array to a 2D string array.
func convertToString(data []int) [][]string {
	array := make([][]string, len(data))
	for i, elem := range data {
		array[i] = []string{strconv.Itoa(elem)}
	}
	return array
}

// Generate a random integer array.
func generateRandomArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(1000000)
	}
	return array
}

// Retrieve file path based on filename.
func getFilePath(filename string) (string, error) {
	paths := map[string]string{
		"output": "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Sort-Data/data.csv",
		"input":  "C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Sort-Data/sorted-data.csv",
	}

	if path, ok := paths[filename]; ok {
		return path, nil
	}
	return "", fmt.Errorf("invalid filename: %s", filename)
}

// Perform quicksort on the array.
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivotIndex := rand.Intn(len(arr))
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	// Generate random array and write to CSV.
	outputArr := generateRandomArray(100)
	outputFormattedArr := convertToString(outputArr)
	writeCSV(outputFormattedArr, "output")

	// Read CSV, build array, and sort.
	data, err := readCSV()
	if err != nil {
		log.Fatal(err)
	}

	unsortedArr := buildArray(data)
	quickSort(unsortedArr)

	// Write sorted array to a new CSV file.
	sortedFormattedArr := convertToString(unsortedArr)
	writeCSV(sortedFormattedArr, "input")
}
