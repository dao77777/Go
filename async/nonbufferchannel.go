package main

import (
	"fmt"
	"time"
)



func main() {
	c := make(chan int)
	fmt.Printf("type of c = %T\n", c)
	go func() {
		defer fmt.Println("goroutine结束")
	
		fmt.Println("goRoutine 正在运行...")
	
		c <- 666
	}()
	num := <-c
	for {
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}

	
}