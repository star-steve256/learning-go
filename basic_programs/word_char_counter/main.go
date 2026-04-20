package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("===== Word and character counter =====")

	path := "file.txt"
	fmt.Println("File path:", path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)

	fmt.Println("Number of bytes:", len(byteData))

	var wordCount = 0
	for i, singleByte := range byteData {
		if string(singleByte) == " " && string(byteData[i+1]) != " " {
			wordCount++
		}
	}
	
	fmt.Println("Number of words:", wordCount)
}
