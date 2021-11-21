package main

import (
	"fmt"
)

func printMap(cityMap map[string]string) {
	for key, val := range cityMap {
		fmt.Println("key = ", key, "val = ", val)
	}
	cityMap["China"] = "Beijing1"
}

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap 是一个空map")
	}
	myMap = make(map[string]string, 10)
	myMap["one"] = "java"
	myMap["two"] = "c++"
	myMap["three"] = "python"
	fmt.Printf("%T\n", myMap)

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "c++"
	myMap2[3] = "python"
	fmt.Printf("%v\n", myMap2)

	myMap3 := map[string]string {
		"one": "php",
		"two": "c++",
		"three": "python",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	cityMap["Japan"] = "wuhu"
	delete(cityMap, "USA")
	printMap(cityMap)
	printMap(cityMap)
}