package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	syncPackage()
}

// 'go' create a subroutine that execute parallel with the main routine
// subroutines only run while main routine are executing
func goRoutine() {
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func syncPackage() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		countWithStop("sheep", 5)
		wg.Done() // decrement count
	}()

	// Wait gonna block main routine until the wg value reach zero (syncronizing subroutines)
	wg.Wait()
}

func countWithStop(thing string, limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

// We can use channels to communicate between subroutines
func channelExample() {
	c := make(chan string)
	go countWithChannel("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countWithChannel(thing string, c chan string) {
	for i := 1; true; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // when done, close the channel to prevent deadlock (only senders should close the channels)
}

func channelWithCapacity() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "World"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

func asyncSubroutinesExecution() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c1 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
