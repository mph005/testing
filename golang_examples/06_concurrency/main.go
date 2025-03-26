package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple function to demonstrate goroutines
func sayHello(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("Hello, %s! (count: %d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

// Function that returns the sum of a slice of integers
func calculateSum(numbers []int, resultChan chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
		// Simulate some work
		time.Sleep(50 * time.Millisecond)
	}
	// Send the result to the channel
	resultChan <- sum
}

// Function that demonstrates waiting for multiple goroutines with WaitGroup
func processData(id int, wg *sync.WaitGroup) {
	// Ensure the WaitGroup counter is decremented when the function returns
	defer wg.Done()

	fmt.Printf("Processing data in goroutine %d\n", id)
	// Simulate processing time
	time.Sleep(time.Duration(id*200) * time.Millisecond)
	fmt.Printf("Finished processing data in goroutine %d\n", id)
}

// Worker function that processes jobs from a channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Simulate work
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back
	}
}

// Function that demonstrates a non-blocking channel send and receive
func nonBlockingExample(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	select {
	case ch <- "hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("No message sent, would block")
	}
}

// Function that demonstrates using a mutex to protect a shared data structure
func mutex() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	// Launch 100 goroutines that increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Lock the mutex before accessing the shared counter
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("Counter value:", counter)
}

func main() {
	fmt.Println("--- Basic Goroutines ---")

	// Start goroutines
	go sayHello("Alice", 3)
	go sayHello("Bob", 3)

	// Main goroutine continues execution
	fmt.Println("Goroutines started")

	// Wait to see the output from goroutines
	// In a real program, you'd use sync.WaitGroup instead
	time.Sleep(1 * time.Second)

	fmt.Println("\n--- Channels ---")

	// Create a channel
	ch := make(chan string)

	// Start a goroutine that sends a message on the channel
	go func() {
		fmt.Println("Sending message on channel...")
		ch <- "Hello from goroutine!"
	}()

	// Receive the message from the channel
	msg := <-ch
	fmt.Println("Received:", msg)

	// Channel as a synchronization mechanism
	done := make(chan bool)

	go func() {
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Done working!")
		done <- true
	}()

	// Wait for work to complete
	<-done
	fmt.Println("Main goroutine continues")

	// Channels with multiple values
	fmt.Println("\n--- Multiple Values on Channels ---")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a channel for results
	resultChan := make(chan int)

	// Split the numbers into two parts and compute the sum in separate goroutines
	go calculateSum(numbers[:len(numbers)/2], resultChan)
	go calculateSum(numbers[len(numbers)/2:], resultChan)

	// Receive both results from the channel
	result1 := <-resultChan
	result2 := <-resultChan

	fmt.Printf("Result 1: %d\n", result1)
	fmt.Printf("Result 2: %d\n", result2)
	fmt.Printf("Total: %d\n", result1+result2)

	// Using WaitGroup to synchronize multiple goroutines
	fmt.Println("\n--- WaitGroup ---")

	var wg sync.WaitGroup

	// Launch 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go processData(i, &wg)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines completed")

	// Worker pool pattern
	fmt.Println("\n--- Worker Pool ---")

	// Create channels for jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel to signal no more jobs

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}

	// Buffered channels
	fmt.Println("\n--- Buffered Channels ---")

	bufferedCh := make(chan string, 2)

	// Send messages without blocking
	bufferedCh <- "Message 1"
	bufferedCh <- "Message 2"

	fmt.Println("Sent 2 messages to buffered channel")

	// Read messages
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)

	// Channel directions (send-only and receive-only)
	fmt.Println("\n--- Channel Directions ---")

	// pingChannel is send-only in pinger and receive-only in printer
	pingChannel := make(chan string)

	// pinger sends, printer receives
	go func(pings chan<- string) { // Send-only channel
		pings <- "ping"
	}(pingChannel)

	go func(pings <-chan string) { // Receive-only channel
		msg := <-pings
		fmt.Println("Received:", msg)
	}(pingChannel)

	time.Sleep(500 * time.Millisecond) // Wait for communication to finish

	// Select statement for multiple channels
	fmt.Println("\n--- Select ---")

	c1 := make(chan string)
	c2 := make(chan string)

	// Send values on each channel
	go func() {
		time.Sleep(200 * time.Millisecond)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		c2 <- "two"
	}()

	// Receive from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received from c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2:", msg2)
		}
	}

	// Timeout with select
	fmt.Println("\n--- Timeout ---")

	timeoutChan := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		timeoutChan <- "response"
	}()

	select {
	case res := <-timeoutChan:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout after 1 second")
	}

	// Non-blocking channel operations with select
	fmt.Println("\n--- Non-blocking Channel Operations ---")

	nonBlockingCh := make(chan string)
	nonBlockingExample(nonBlockingCh)

	// Demonstrate mutex for shared state
	fmt.Println("\n--- Mutex ---")
	mutex()
}
