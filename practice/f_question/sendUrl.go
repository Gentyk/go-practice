package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sender(url string, client *http.Client, chan_in chan<- string) error {
	_, err := client.Get(url)
	s := url + " url - "
	if err != nil {
		fmt.Println("ok")
		chan_in <- s + "ok"
		return nil
	}
	fmt.Println(s + "not ok")
	chan_in <- s + "not ok"
	return err
}

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	var wg sync.WaitGroup

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	chan1 := make(chan string)

	for _, i := range urls {
		wg.Add(1)
		go func(url string, client *http.Client, chan_in chan<- string, wg *sync.WaitGroup) {
			defer wg.Done()
			sender(url, client, chan1)
		}(i, &client, chan1, &wg)
	}

	go func() {
		wg.Wait()
		close(chan1)
	}()

	for i := range chan1 {
		fmt.Println(i)
	}
}
