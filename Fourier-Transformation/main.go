package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
	"runtime"

	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/mjibson/go-dsp/fft"
)

const N = 4096

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

func main() {
	path := getFilePath()
	f := readAudioFile(path)
	defer f.Close()

	streamer, format := decodeAudio(f)
	defer streamer.Close()

	samples := readSample(streamer)
	fourier := Dft(samples)

	fourier2 := fft.FFTReal(samples)

	// Use the Dft result directly
	fmt.Printf("%T\n", streamer)
	fmt.Printf("%T\n", format)
	fmt.Println(fourier[0])
	fmt.Println(fourier2[0])
}