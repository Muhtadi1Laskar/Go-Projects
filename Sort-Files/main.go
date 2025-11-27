package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}

type SortBy int

const (
	NameAsc SortBy = iota
	NameDesc
	SizeAsc
	SizeDesc
	TimeAsc
	TimeDesc
)

func LoadFiles(path string) ([]FileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	files := make([]FileInfo, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}

		files = append(files, FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
		})
	}

	return files, nil
}

func SortFiles(files []FileInfo, mode SortBy) {
	sort.Slice(files, func(i, j int) bool {
		switch mode {
		case NameAsc:
			return files[i].Name < files[j].Name
		case NameDesc:
			return files[i].Name > files[j].Name
		case SizeAsc:
			return files[i].Size < files[j].Size
		case SizeDesc:
			return files[i].Size > files[j].Size
		case TimeAsc:
			return files[i].ModTime.Before(files[j].ModTime)
		case TimeDesc:
			return files[i].ModTime.After(files[j].ModTime)
		default:
			return true
		}
	})
}

func PrintFiles(files []FileInfo) {
	for _, f := range files {
		fmt.Printf("Name: %-20s  Size: %-10d  Modified: %s\n",
			f.Name, f.Size, f.ModTime.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	path := filepath.Join("C:/Users/laska/OneDrive/Documents/Coding/Golang/Go-Projects/Sort-Files/files")

	files, err := LoadFiles(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n📁 Files sorted by Name (ASC):")
	SortFiles(files, NameAsc)
	PrintFiles(files)

	fmt.Println("\n📁 Files sorted by Size (DESC):")
	SortFiles(files, SizeDesc)
	PrintFiles(files)

	fmt.Println("\n📁 Files sorted by Modified Time (ASC):")
	SortFiles(files, TimeAsc)
	PrintFiles(files)
}
