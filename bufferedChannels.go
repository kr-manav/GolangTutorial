package main

import (
	"fmt"
	"time"
)

// func fibonacci(n int, ch chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		ch <- x
// 		x, y = y, x+y
// 	}
// 	close(ch)
// }

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)
	fmt.Println(len(ch))
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
	time.Sleep(5000 * time.Millisecond)

	// ch := make(chan int, 10)
	// go fibonacci(cap(ch), ch)
	// for i := range ch {
	// 	fmt.Println(i)
	// }

}
