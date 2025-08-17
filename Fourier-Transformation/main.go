package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
	"runtime"
	"sort"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
)

const N = 4096

type freqVal struct {
	freq float64
	mag float64
}

func getFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "./sounds/sample.wav")
}

func Dft(signal []float64) []complex128 {
	var N int = len(signal)
	var result = make([]complex128, N)

	for k := 0; k < N; k++ {
		var s complex128 = 0
		for n := 0; n < N; n++ {
			angle := -2 * math.Pi * float64(k) * float64(n) / float64(N)
			s += complex(float64(signal[n]), 0) * cmplx.Exp(complex(0, angle))
		}
		result[k] = s
	}

	return result
}

func readAudioFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return f
}

func decodeAudio(f *os.File) (beep.StreamSeekCloser, beep.Format) {
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return streamer, format
}

func readSample(streamer beep.StreamSeekCloser) []float64 {
	samples := make([]float64, N)
	buf := make([][2]float64, 1)
	for i := 0; i < N; i++ {
		sample, ok := streamer.Stream(buf)
		if !ok || sample == 0 {
			break
		}
		samples[i] = buf[0][0]
	}
	return samples
}

func convertToMagnitude(complexResult []complex128, format beep.Format) []freqVal {
	var result []freqVal

	for i := 0; i < len(complexResult); i++ {
		re := real(complexResult[i])
		im := imag(complexResult[i])
		mag := math.Sqrt(re * re + im * im)
		freq := float64(i) * float64(format.SampleRate) / float64(N)
		result = append(result, freqVal{freq, mag})
	}
	return result
}

func main() {
	path := getFilePath()
	f := readAudioFile(path)
	defer f.Close()

	streamer, format := decodeAudio(f)
	defer streamer.Close()

	samples := readSample(streamer)
	fourier := Dft(samples)

	result := convertToMagnitude(fourier, format)

	sort.Slice(result, func(i, j int) bool {
		return result[i].mag > result[j].mag
	})

	for i := 0; i < 50; i++ {
		fmt.Printf("%d. %.1f Hz (mag=%.2f)\n", i+1, result[i].freq, result[i].mag)
	}
}