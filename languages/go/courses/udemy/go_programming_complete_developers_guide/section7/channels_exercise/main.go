package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job int

func longCalculation(i Job) int {
	duration := time.Duration(rand.Intn(1000)+100) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", i, duration)
	return int(i) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 500)

	for i := 0; i < 500; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}

	return jobs
}

func runJob(resultChan chan int, i Job) {
	resultChan <- longCalculation(i)
}

func main() {
	fmt.Printf("\nChannels Exercise.\n")
	jobs := makeJobs()

	resultChan := make(chan int, 10)

	for _, v := range jobs {
		go runJob(resultChan, v)
	}

	resultCount := 0
	sum := 0

	for {
		result := <-resultChan
		sum += result
		resultCount += 1

		if resultCount == len(jobs) {
			break
		}
	}

	fmt.Println("Sum is:", sum)
}
