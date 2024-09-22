package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readData() ([][]string, error) {
	fileName, err := os.Open("data.csv");
	if err != nil {
		return nil, err;
	}
	defer fileName.Close();

	rawData := csv.NewReader(fileName);
	data, err :=rawData.ReadAll();
	if err != nil {
		return nil, err;
	}

	return data, nil
}

func buildArray(data [][]string) []int {
	var array []int = []int{};
	for _, elem := range data {
		num, _ := strconv.Atoi(elem[0]);
		array = append(array, num);
	} 
	return array;
}

func insertionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		var key int = array[i];
		var j int = i - 1;

		for j >= 0 && array[j] > key {
			array[j+1] = array[j];
			j--;
		}
		array[j+1] = key;
	}
	return array;
}

func main() {
	data, err := readData();
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}

	var array []int = buildArray(data);
	var sortedArray []int = insertionSort(array);

	fmt.Println(sortedArray);
}