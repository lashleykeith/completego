package main

import (
	"fmt"
)

func main() {


	//Int
	colors := make(map[int]string)
	colors[10] = "#ffffff"
	delete(colors, 10)
	fmt.Println(colors)
//////////////////////////////////////////////////
	// Strings
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#00ff00",

	// }
	
	// colors["white"] = "#ffffff"


	//fmt.Println(colors)
}

