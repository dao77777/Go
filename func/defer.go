package main

import (
	"fmt"
)

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	defer func1()
	defer func2()
	defer func3()

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}