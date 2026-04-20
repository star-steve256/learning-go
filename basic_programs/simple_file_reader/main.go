package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("=====Simple file reader =====")

	path := "file.txt"
	file, err:= os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	fmt.Print(string(byteData))
}
