package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

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

func DecodeAudio(filepath string, start, end int) []float64 {
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

	var startTime int = 175
	var endTime int = 180
	var samples []float64 = DecodeAudio(filePath, startTime, endTime)

	WriteWavFile(outPutFile, samples)

}
