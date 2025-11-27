package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashFile(path string) ([32]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return [32]byte{}, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return [32]byte{}, err
	}

	var sum [32]byte
	copy(sum[:], h.Sum(nil))
	return sum, nil
}

func main() {
	var filePathString string = "C:/Users/laska/Downloads/Pictures/JPG Format/IMG_3291 12x16.jpg"
	fmt.Println(hashFile(filePathString))

}