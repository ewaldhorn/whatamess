package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Hits struct {
	count int
	sync.Mutex
}

func wait() {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()

	hits.count += 1
	fmt.Println("Served iteration", iteration)
}

func doDemo() {
	var wg sync.WaitGroup
	hitCounter := Hits{}

	for i := 1; i <= 200; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Println("Waiting for all the threads to complete.")
	wg.Wait()
	fmt.Println("All done.")

	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()

	fmt.Println()
	fmt.Println("Total hits were", totalHits)
}
