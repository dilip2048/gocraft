// this is a sample example of goroutine and channel in go

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var channel chan string
var waitGroup sync.WaitGroup

// The main function will initialize the chan and starts the goroutine function
func main()  {
	channel = make(chan string)
	for i := 0; i < 3; i++ {
		// increase the counter of waitgroup to track the goroutines
		waitGroup.Add(1)
		// this will start 3 go routines
		go worker()
	}

	// add data to the channel
	for i:= 0; i < 5; i++ {
		channel <- strconv.Itoa(i)
	}

	close(channel)
	// waits until all the goroutines are done
	waitGroup.Wait()
}

// The worker function represent every goroutine
func worker() {
	fmt.Println(fmt.Sprintf("Started go routine: %v", time.Now()))
	defer func() {
		// after the method is over, send signal to main method that this goroutine done its work
		fmt.Println("Terminating goroutine")
		waitGroup.Done()
	}()
	for {
		value, ok := <-channel
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println("The value in the channel is: " + value)
		time.Sleep(time.Second)
	}
}
