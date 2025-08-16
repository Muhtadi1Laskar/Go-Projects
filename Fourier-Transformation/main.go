package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func dft(signal []float32) []complex128 {
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

func main() {
	var samples []float32 = []float32{1.0, 0.0, -1.0, 0.0}
	spectrum  := dft(samples)

	fmt.Println(spectrum)
}