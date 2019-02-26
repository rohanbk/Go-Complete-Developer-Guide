package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com/",
		"http://www.twitter.com/",
		"http://www.amazon.com/",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLinkStatus(c, link)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLinkStatus(c, link)
		}(l)
	}
}

func checkLinkStatus(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("There was an issue connecting with %v, http%v\n", link, err)
		c <- link
		return
	}
	fmt.Printf("There was no issue connecting with %v\n", link)
	c <- link

}
