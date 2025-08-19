package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/faiface/beep/mp3"
)

func getFilePath(folderName string) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, folderName)
}

func readAudioFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func EncodeAudio(filepath string, start, end int) []float64 {
	f := readAudioFile(filepath)
	defer f.Close()

	streamer, _, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	var samples []float64
	buf := make([][2]float64, 1)
	for {
		sample, ok := streamer.Stream(buf)
		if !ok || sample == 0 {
			break
		}
		samples = append(samples, float64(buf[0][0]))
	}
	samplesPerSecond := 44100
	startIndex := start * samplesPerSecond
	endIndex := end * samplesPerSecond
	cutSamples := samples[startIndex:endIndex]

	return cutSamples
}

func main() {
	var filePath string = getFilePath("./Audio/sample.mp3")
	var startTime int = 10
	var endTime int = 70
	samples := EncodeAudio(filePath, startTime, endTime)

	fmt.Println(len(samples))
}
