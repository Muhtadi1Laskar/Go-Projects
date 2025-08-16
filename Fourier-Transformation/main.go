package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
	"runtime"

	"github.com/faiface/beep/wav"
)

func Dft(signal []float32) []complex128 {
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

func getFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "./sounds/sample.wav")
}

func main() {
	path := getFilePath()
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	fmt.Println(streamer)
	fmt.Println(format)
}