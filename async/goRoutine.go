package main

import (
	"fmt"
	"time"
	"runtime"
)

func goRoutine1() {
	i := 0
	for {
		if i > 10 {
			runtime.Goexit()
		}
		fmt.Println("goRoutine run", i)
		time.Sleep(1 * time.Second)
		i++
	}
}

func goRoutine2(a int, b int) {
	fmt.Println("a =", a, "b =", b)
}

func main() {
	fmt.Printf("type of goRoutine = %T\n", goRoutine1)
	closure := func() (func()) {
		wow := 1
		return func() {
			fmt.Println(wow)
		}
	}
	closure()()
	go goRoutine1()
	i := 0
	for {
		fmt.Println("mainRoutine run", i)
		time.Sleep(1 * time.Second)
		i++
	}

}