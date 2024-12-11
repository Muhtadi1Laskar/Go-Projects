package main

import (
	"crypto/sha256"
	// "crypto/md5"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Identicon struct {
	Hex       []byte
	Color     color.RGBA
	Grid      []Cell
	PixelMap  []Pixel
	ImageSize int
	CellSize  int
}

type Cell struct {
	Value int
	Index int
}

type Pixel struct {
	TopLeft     image.Point
	BottomRight image.Point
}

func hash(data string) []byte {
	hash := sha256.New()
	hash.Write([]byte(data))

	return hash.Sum(nil)
}

func HashInput(input string) Identicon {
	hash := hash(input)
	return Identicon{
		Hex:       hash[:],
		ImageSize: 350,
		CellSize:  50,
	}
}

func PickColor(identicon Identicon) Identicon {
	identicon.Color = color.RGBA{
		R: identicon.Hex[0],
		G: identicon.Hex[1],
		B: identicon.Hex[2],
		A: 255,
	}
	return identicon
}

func BuildGrid(identicon Identicon) Identicon {
	var grid []Cell
	for i := 0; i < 15; i += 3 {
		row := []int{int(identicon.Hex[i]), int(identicon.Hex[i+1]), int(identicon.Hex[i+2])}
		row = append(row, row[1], row[0])
		for _, value := range row {
			grid = append(grid, Cell{Value: value, Index: len(grid)})
		}
	}
	identicon.Grid = grid
	return identicon
}

func FilterOddSquares(identicon Identicon) Identicon {
	var filteredGrid []Cell
	for _, cell := range identicon.Grid {
		if cell.Value%2 == 0 { 
			filteredGrid = append(filteredGrid, cell)
		}
	}
	identicon.Grid = filteredGrid
	return identicon
}

func BuildPixelMap(identicon Identicon) Identicon {
	var pixelMap []Pixel
	for _, cell := range identicon.Grid {
		x := (cell.Index % 5) * identicon.CellSize
		y := (cell.Index / 5) * identicon.CellSize
		topLeft := image.Point{x, y}
		bottomRight := image.Point{x + identicon.CellSize, y + identicon.CellSize}
		pixelMap = append(pixelMap, Pixel{TopLeft: topLeft, BottomRight: bottomRight})
	}
	identicon.PixelMap = pixelMap
	return identicon
}

func DrawImage(identicon Identicon, input string) {
	img := image.NewRGBA(image.Rect(0, 0, identicon.ImageSize, identicon.ImageSize))

	for _, pixel := range identicon.PixelMap {
		for y := pixel.TopLeft.Y; y < pixel.BottomRight.Y; y++ {
			for x := pixel.TopLeft.X; x < pixel.BottomRight.X; x++ {
				img.Set(x, y, identicon.Color)
			}
		}
	}

	fileName := fmt.Sprintf("%s.png", input)
	file, err := os.Create("Identicon-Generator/image.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}
	fmt.Printf("Identicon saved as %s\n", fileName)
}


func main() {
	var input string
	fmt.Print("Enter a string to generate Identicon: ")
	fmt.Scanln(&input)

	identicon := HashInput(input)
	identicon = PickColor(identicon)
	identicon = BuildGrid(identicon)
	identicon = FilterOddSquares(identicon)
	identicon = BuildPixelMap(identicon)
	DrawImage(identicon, input)
}
