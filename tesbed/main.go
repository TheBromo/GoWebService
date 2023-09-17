package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Millisecond)

	for {
		select {
		//TODO doesnt work yet
		case <-ticker.C:
			println("timed")
		default:
			println("constant")
		}

	}
}
