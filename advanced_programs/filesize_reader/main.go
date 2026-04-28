package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	var inputPath string

	fmt.Print("Enter file path: ")
	fmt.Scan(&inputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteSize, err := countFileByteSize(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	printByteSizeFormats(byteSize)
}

func countFileByteSize(file *os.File) (size int, err error) {
	byteData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	return len(byteData), nil
}

func printByteSizeFormats(size int) {
	byteSize := float64(size)
	kiloBytes := byteSize / 1024
	megaBytes := byteSize / (math.Pow(1024, 2))
	gigaBytes := byteSize / (math.Pow(1024, 3))

	fmt.Println("File sizes:")
	fmt.Printf("%.2fKB; %.2fMB; %.2fGB\n", kiloBytes, megaBytes, gigaBytes)
}
