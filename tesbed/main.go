package main

import (
	"time"
)

var (
	listeners = make([]chan string, 0)
)

func main() {
	test1 := make(chan string, 0)
	test2 := make(chan string, 0)
	test3 := make(chan string, 0)

	listeners = append(listeners, test1, test2, test3)

	go func() {
		for {
			listeners[0] <- "test 1"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			listeners[1] <- "test 2"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			listeners[2] <- "test 3"
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		v1 := <-test1
		println(v1)

		v2 := <-test2
		println(v2)

		v3 := <-test3
		println(v3)
	}
}
