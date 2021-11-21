package main

import (
	"fmt"
	// "time"
)

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <- quit:
			fmt.Println("quit, x =", x, "y =", y)
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("c =", <-c)
		}
		quit <- 0
	}()
	// time.Sleep(2 * time.Second)
	// for i := 0; i < 6; i++ {
	// 	c <- i
	// }
	// time.Sleep(2 * time.Second)
	// fmt.Println("quit =", <-quit)
	fibonacii(c, quit)
}