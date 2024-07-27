package main

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
	}
	c := make(chan string, len(links))
	var regex = regexp.MustCompile(` [a-z]+|[A-Z]+`)
	var valid = regexp.MustCompile(`up|UP`)
	channels := 10
	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, link := range links {
		wg.Add(1)
		go checkLink(link, c, &wg, &channels, &lock)
	}

	for channels > 0 {
		wg.Add(1)
		go func(l string, c chan string) {
			time.Sleep(2 * time.Second)
			if valid.MatchString(l) {
				l = regex.ReplaceAllString(l, "")
				fmt.Println(l, "It's working", channels, "Routine")
			} else {
				l = regex.ReplaceAllString(l, "")
				fmt.Println(l, "It's off", channels, "Routine")
			}
			checkLink(l, c, &wg, &channels, &lock)
		}(<-c, c)
	}
	fmt.Println("End Main Go routine")
	wg.Wait()
	fmt.Println("End Wait for all routines")
}

func checkLink(link string, c chan<- string, wg *sync.WaitGroup, limit *int, l *sync.Mutex) {
	if _, err := http.Get(link); err != nil {
		c <- fmt.Sprint(link, " off")
		return
	}
	c <- fmt.Sprint(link, " up")
	wg.Done()
	l.Lock()
	*limit--
	l.Unlock()
}
