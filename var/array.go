package main

import (
	"fmt"
)

func printArray(myArray [10]int) {
	for index, val := range myArray {
		fmt.Println("index = ", index, ", val = ", val)
	}
}

func main() {
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}
	for index, val := range myArray2 {
		fmt.Println("index =", index, "val =", val)
	}

	printArray(myArray1)
}