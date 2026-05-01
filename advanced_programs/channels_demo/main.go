package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	fmt.Println("Channels demo...")

	var wg sync.WaitGroup

	mainChannel := make(chan [2]int)
	count := 2
	wg.Add(count)

	for i := range count {
		go addData(i, &mainChannel, &wg)
	}

	// I don't understand this point (need to clarify)
	go func ()  {
		wg.Wait()
		close(mainChannel)
	}()

	for dataArr := range mainChannel {
		fmt.Printf("Go routine %d says: %d\n", dataArr[0], dataArr[1])
	}

}

func addData(index int, channel *chan [2]int, wg *sync.WaitGroup) {
	defer wg.Done()
	randInt := rand.Int()
	dataArr := [2]int{index, randInt}

	*channel <- dataArr
}
