package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	//readChannel(c)
	for i := 0; i < len(links); i++ {
		fmt.Print(<-c)
	}
}

func checkLink(link string, c chan<- string) {
	if _, err := http.Get(link); err != nil {
		c <- fmt.Sprintln(link, "off")
		return
	}
	c <- fmt.Sprintln(link, "up")

}

func readChannel[V string | int32](c <-chan V) {
	for i := 0; i < len(c); i++ {
		fmt.Print(<-c)
	}
}
