package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	go sender(0, c)
	go sender(1, c)
	go receiver(c)
	select {}
}

func sender(thread int, c chan int) {
	for {
		r := rand.Intn(100)
		if r%5 == 0 {
			fmt.Println("Thread", thread, "sending", r)
			c <- r
		} else {
			time.Sleep(5 * time.Second)
		}
	}
}

func receiver(c chan int) {
	for {
		select {
		case r := <-c:
			{
				fmt.Println("Received:", r)
			}
		default:
		}
	}
}
