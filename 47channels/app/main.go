package main

import (
		"fmt"
		"net/http"
        "time"
		)

func main(){
        links := []string{
                "http://google.com",
                "http://facebook.com",
                "http://stackoverflow.com",
                "http://golang.com",
                "http://amazon.com",
        }
        
        c := make(chan string)

        for _, link := range links {
                go checkLink(link, c)
        }

        for l := range c {
           go func(link string){
            time.Sleep(5 * time.Second)
            checkLink(link,c)
           }(l)
        }
    }

func checkLink(link string, c chan string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
        c <- link
		return
	}

	fmt.Println(link, " is up!")
        c <- link
}

//https://golang.org/pkg/time/#Sleep

/*
seconds := 10
fmt.Print(time.Duration(seconds)*time.Second) // prints 10s

const (
        Nanosecond  Duration = 1
        Microsecond          = 1000 * Nanosecond
        Millisecond          = 1000 * Microsecond
        Second               = 1000 * Millisecond
        Minute               = 60 * Second
        Hour                 = 60 * Minute
)
*/