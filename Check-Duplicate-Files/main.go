package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func findDuplicateFiles(root string) (map[[32]byte][]string, error) {
	sizeGroups := make(map[int64][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		sizeGroups[info.Size()] = append(sizeGroups[info.Size()], path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	duplicates := make(map[[32]byte][]string)

	for _, files := range sizeGroups {
		if len(files) < 2 {
			continue
		}
		for _, file := range files {
			hash, err := hashFile(file)
			if err != nil {
				return nil, err
			}
			duplicates[hash] = append(duplicates[hash], file)
		}
	}

	result := make(map[[32]byte][]string)
	for hash, paths := range duplicates {
		if len(paths) > 1 {
			result[hash] = paths
		}
	}

	return result, nil
}

func main() {
	root := "C:/Users/laska/Videos"

	dups, err := findDuplicateFiles(root)
	if err != nil {
		panic(err)
	}

	fmt.Println("Duplicate Files:")
	for hash, files := range dups {
		fmt.Printf("\nHash: %x\n", hash)
		for _, f := range files {
			fmt.Println("  -", f)
		}
	}

}
