package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func channelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	outChan := make(chan int)
	errChan := make(chan error)
	finishChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(outChan chan<- int, errChan chan<- error, val int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
			fmt.Println("finished job id:", val)
			outChan <- val
			if val == 60 {
				errChan <- errors.New("error in 60")
			}

		}(outChan, errChan, i, &wg)
	}

	go func() {
		wg.Wait()
		fmt.Println("finish all job")
		close(finishChan)
	}()

Loop:
	for {
		select {
		case val := <-outChan:
			fmt.Println("finished:", val)
		case err := <-errChan:
			fmt.Println("error:", err)
			break Loop
		case <-finishChan:
			break Loop
		case <-time.After(100000 * time.Millisecond):
			break Loop
		}
	}

	channelSelect()
}
