package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func InitConcurrency() {
	var totalCount int64 = 0
	var totalTime int64 = 0

	var wg sync.WaitGroup

	for range 10 {
		wg.Go(func() {addToCount(&totalCount, &totalTime)})
	}

	wg.Wait()

	fmt.Println("Total count:", totalCount)
	fmt.Println("Total time of all routines:", totalTime, "ms")
}

func addToCount(total *int64, totalTime *int64) {
	start := time.Now()
	atomic.AddInt64(total, 1)
	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)
	atomic.AddInt64(totalTime, elapsed.Milliseconds())
}
