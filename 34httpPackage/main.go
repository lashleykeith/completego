package main

import (
		"fmt"
		"net/http"
		"os"
		)



func main() {
	resp, err := http.Get("http://google.com")
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

// https://golang.org/pkg/
// https://golang.org/pkg/net/http/
// https://golang.org/pkg/net/http/#Get