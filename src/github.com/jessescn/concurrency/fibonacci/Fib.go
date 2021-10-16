package main

import "fmt"

// calculate fib values with subroutines communcating using channels
// this approach optimize the results, however do not guarantee the correct output order
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// we can create a bunch of workers that will calculate together the results
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
}

// creating a worker that receives a channel with numbers to calculate (jobs) and channel to put its results
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
