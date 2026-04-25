package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var finishedApiProcessCount int64 = 0

func main() {
	// uses only one core for true concurrency

	runtime.GOMAXPROCS(1)
	start := time.Now()
	var wg sync.WaitGroup

	var routineCount = 10
	for i := 0; i < routineCount; i++ {
		wg.Go(func() {proceedRoutine(i + 1)})
	}
	
	wg.Wait()
	fmt.Println("Number of API calls finished:", finishedApiProcessCount)

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %fs\n", elapsed.Seconds())
}

type ApiResponse struct {
	Fact string `json:"fact"`
}

func proceedRoutine(routineIndex int) {
	fact, err := requestApi(routineIndex)
	if err != nil {
		fmt.Printf("ERROR at routine %d: %s\n", routineIndex, err)
		return
	}

	printFact(routineIndex, fact)
	atomic.AddInt64(&finishedApiProcessCount, 1)
}

func requestApi(routineIndex int) (fact string, err error) {
	var apiResponse ApiResponse
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Printf("API call %d failed: %s\n", routineIndex, err)
		return "", err
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("IO read %d failed: %s\n", routineIndex, err)
		return "", err
	}

	err = json.Unmarshal(byteData, &apiResponse)
	if err != nil {
		fmt.Printf("JSON decode %d failed: %s\n", routineIndex, err)
		return "", err
	}

	return apiResponse.Fact, nil
}

func printFact(routineIndex int, fact string) {
	fmt.Printf("Fact from routine %d: %s\n", routineIndex, fact)
}
