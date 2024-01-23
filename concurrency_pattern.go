package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// performTask simulates a task and sends the result on the channel.
func performTask(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup when the task is completed

	// Simulate a task
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	ch <- fmt.Sprintf("Task %d completed", id)
}

// startTasks initiates the specified number of tasks as goroutines and waits for their completion.
func startTasks(tasks int, ch chan<- string) {
	var wg sync.WaitGroup

	// Start goroutines
	for i := 1; i <= tasks; i++ {
		wg.Add(1)
		go performTask(i, ch, &wg)
	}

	// Use a goroutine to close the channel when all tasks are completed
	go func() {
		wg.Wait()
		close(ch)
	}()
}

// receiveResults receives and prints results from the channel.
func receiveResults(ch <-chan string) {
	// Receive data from the channel
	for result := range ch {
		fmt.Println(result)
	}
}

func main() {
	tasks := 5
	ch := make(chan string)

	// Start goroutines for tasks and close channel when completed
	startTasks(tasks, ch)

	// Receive and print results
	receiveResults(ch)
}
