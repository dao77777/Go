package main

import (
	"fmt"
)

func printArray(myArray []int) {
	for index, val := range myArray {
		fmt.Println("index = ", index, "val = ", val)
	}
	myArray[0] = 2
	return
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组, 切片 slice
	fmt.Printf("%T\n", myArray)
	printArray(myArray)
	for index, val := range myArray {
		fmt.Println("index = ", index, "val = ", val)
	}

	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	var slice2 []int
	slice2 = make([]int, 3)
	slice2[0] = 2
	fmt.Printf("slice2 = %v\n", slice2)

	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	slice3 := numbers[0:3]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice3), cap(slice3), slice3)
}