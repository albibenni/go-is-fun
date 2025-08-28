package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 9; r++ {
		r, ok := <-results
		if !ok {
			fmt.Printf("Not ok\n")
		} else {
			fmt.Printf("Results: %d\n", r)
		}
	}
}

// select {
// case msg1 := <-ch1:
//     fmt.Println("Received from ch1:", msg1)
// case msg2 := <-ch2:
//     fmt.Println("Received from ch2:", msg2)
// case <-time.After(1 * time.Second):
//     fmt.Println("Timeout!")
// default:
//     fmt.Println("No channels ready")
// }
