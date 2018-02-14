package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://kickdudes.club",
		"http://kickdudes.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://kapook.com",
		"http://twitter.com",
		"http://twinsynergy.co.th",
	}

	c := make(chan string)

	for _, url := range links {
		go checkMyURL(url, c)
	}
	for range links {
		fmt.Println(<-c)
	}

}

func checkMyURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		//fmt.Println(url, " might be down! ", err)
		c <- url + " Might be down! "
		fmt.Println(err)
		return
	}
	//fmt.Println(url, " status is up.")
	c <- url + " status is up."
}
