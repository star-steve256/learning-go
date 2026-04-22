package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("===== File writer =====")

	var fileName string
	fmt.Print("Enter txt file name: ")
	_, err := fmt.Scan(&fileName)
	if err != nil {
		log.Fatal("Error reading file name:", err)
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	textContent, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading text content:", err)
	}

	_, err = file.Write([]byte(textContent))
	if err != nil {
		log.Fatal("Error writing file:", err)
	}

	fmt.Println("File writing complete.")
}
