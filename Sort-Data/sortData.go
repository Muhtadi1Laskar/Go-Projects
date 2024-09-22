package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readCSV() ([][]string, error) {
	fileName, err := os.Open("data.csv");
	if err != nil {
		return nil, err;
	}
	defer fileName.Close();

	rawData := csv.NewReader(fileName);
	data, err := rawData.ReadAll();
	if err != nil {
		return nil, err;
	}

	return data, nil
}

func writeCSV(data [][]string) {
	fileName, err := os.Create("sorted-data.csv");
	if err != nil {
		fmt.Println("Cannot read the CSV file");
		os.Exit(1);
	}
	defer fileName.Close();

	csvWrite:= csv.NewWriter(fileName);
	if err = csvWrite.WriteAll(data); err != nil {
		fmt.Println("Cannot read the CSV file");
		os.Exit(1);
	}
	
	csvWrite.Flush();
	if err := csvWrite.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
		os.Exit(1)
	}
}

func buildArray(data [][]string) []int {
	array := make([]int, 0, len(data));
	for _, elem := range data {
		num, _ := strconv.Atoi(elem[0]);
		array = append(array, num);
	} 
	return array;
}

func converToString(data []int) [][]string {
	array := make([][]string, len(data));
	for i, elem := range data {
		array[i] = []string{strconv.Itoa(elem)};
	}
	return array;
}

func insertionSort(array []int) [][]string {
	for i := 0; i < len(array); i++ {
		var key int = array[i];
		var j int = i - 1;

		for j >= 0 && array[j] > key {
			array[j+1] = array[j];
			j--;
		}
		array[j+1] = key;
	}
	return converToString(array);
}

func main() {
	data, err := readCSV();
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}

	var array []int = buildArray(data);
	var sortedArray [][]string = insertionSort(array);

	writeCSV(sortedArray);

	fmt.Println(sortedArray);
}