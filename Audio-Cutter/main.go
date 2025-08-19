package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

const SAMPLE_RATE = 44100

func getFilePath(folderName string) string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, folderName)
}

func timeStringToSecond(timeStr string) (int, error) {
	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return 0, err
	}
	return t.Hour()*3600 + t.Minute()*60 + t.Second(), nil
}

func DecodeAudioAndCutSample(filepath string, start, end int) []float64 {
	if start > end {
		log.Fatal("Start cannot be greater than the End")
	}
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
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
	startIndex := start * SAMPLE_RATE
	endIndex := end * SAMPLE_RATE
	cutSamples := samples[startIndex:endIndex]

	return cutSamples
}

func WriteWavFile(outPath string, cutSample []float64) {
	f, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enc := wav.NewEncoder(f, SAMPLE_RATE, 16, 1, 1)

	ints := make([]int, len(cutSample))
	for i, v := range cutSample {
		ints[i] = int(v * 32767.0)
	}

	buf := &audio.IntBuffer{
		Data:           ints,
		Format:         &audio.Format{NumChannels: 1, SampleRate: SAMPLE_RATE},
		SourceBitDepth: 16,
	}

	if err := enc.Write(buf); err != nil {
		log.Fatal(err)
	}

	enc.Close()
	fmt.Println("Successfully edited the audio")
}

func main() {
	var filePath string = getFilePath("./Audio/sample.mp3")
	var outPutFile string = getFilePath("./Output/output.wav")
	var startTimeStr string = "00:01:37"
	var endTimeStr string = "00:02:04"

	startTime, err := timeStringToSecond(startTimeStr)
	if err != nil {
		log.Fatal("Invalid start time: ", err)
	}

	endTime, err := timeStringToSecond(endTimeStr)
	if err != nil {
		log.Fatal("Invalid end time: ", err)
	}

	var samples []float64 = DecodeAudioAndCutSample(filePath, startTime, endTime)

	WriteWavFile(outPutFile, samples)
}
