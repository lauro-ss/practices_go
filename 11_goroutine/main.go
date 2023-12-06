package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}
	c := make(chan string)
	var regex = regexp.MustCompile(` [a-z]+|[A-Z]+`)
	var valid = regexp.MustCompile(`up|UP`)
	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		go func(l string, c chan string) {
			time.Sleep(2 * time.Second)
			if valid.MatchString(l) {
				l = regex.ReplaceAllString(l, "")
				fmt.Println(l, "It's working")
			} else {
				l = regex.ReplaceAllString(l, "")
				fmt.Println(l, "It's off")
			}
			checkLink(l, c)
		}(<-c, c)
	}
}

func checkLink(link string, c chan<- string) {
	if _, err := http.Get(link); err != nil {
		c <- fmt.Sprint(link, " off")
		return
	}
	c <- fmt.Sprint(link, " up")
}
