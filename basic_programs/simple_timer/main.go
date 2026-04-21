package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("===== Simple timer =====")

	var min int
	var sec int

	fmt.Println("Set timer: ")
	fmt.Print("Min: ")
	_, err := fmt.Scan(&min)
	if err != nil {
		log.Fatal("Error scanning minute:", err)
	}
	fmt.Print("Sec: ")
	_, err = fmt.Scan(&sec)
	if err != nil {
		log.Fatal("Error scanning second:", err)
	}

	var totalSeconds = (min * 60) + sec
	for i := totalSeconds; i >= 0; i-- {
		var min = i / 60
		var sec = i % 60
		fmt.Printf("\rTime remaining: %02d:%02d", min, sec)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Time's up!")
}
