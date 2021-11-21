package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title		string		`json: "title"`
	Year		int				`json: "year"`
	Price		int				`json: "rmb`
	Actors	[]string	`json: "actors"`	
}

func main() {
	movie := Movie{
		Title:	"喜剧之王",
		Year:		2000,
		Price:	10,
		Actors:	[]string{"xingye", "zhangbozhi"},
	}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
	}
	fmt.Println(myMovie)
}