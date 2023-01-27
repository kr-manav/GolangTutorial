package main

import (
	"fmt"
)

func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y

		// case <-quit:
		// 	fmt.Println("Break")
		// 	return
		// }

		// default:
		// 	time.Sleep(100 * time.Millisecond)
		// 	fmt.Println("Break")
		// 	return

		default:
			fmt.Println("Break")
			return

		}

	}
}

func main() {
	ch := make(chan int, 10)
	quit := make(chan int, 1)

	go func() {
		for i := 0; i < cap(ch); i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(ch, quit)

}
