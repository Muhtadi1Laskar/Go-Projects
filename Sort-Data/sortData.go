package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func readCSV() ([][]string, error) {
	fileName, err := os.Open("C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Sort-Data/data.csv");
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
	fileName, err := os.Create("C:/Users/SYSNET/OneDrive/Documents/Coding/Golang/projects/Sort-Data/sorted-data.csv");
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

func generateRandomArray() [][]string {
	var size int = 100000;
	var array []int = make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(1000000)
	}
	return converToString(array)
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

func merge(left, right []int) []int {
	var merged []int
	var i int = 0
	var j int = 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
		} else {
			merged = append(merged, right[j])
		}
	}
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	var mid int = len(array) / 2
	var left []int = mergeSort(array[:mid])
	var right []int = mergeSort(array[mid:])

	return merge(left, right)
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Pick a pivot index randomly
	pivotIndex := rand.Intn(len(arr))

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Partitioning process
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Move pivot to its final place
	arr[left], arr[right] = arr[right], arr[left]

	// Sort the left and right subarrays
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

func main() {
	data, err := readCSV();
	if err != nil {
		fmt.Println(err);
		os.Exit(1);
	}

	var array []int = buildArray(data);
	// var sortedArray [][]string = insertionSort(array);

	QuickSort(array);
	var stringArray [][]string = converToString(array)
	writeCSV(stringArray);

	// fmt.Println(sortedArray);

	// var randomNum [][]string = generateRandomArray()

	// writeCSV(randomNum)
}