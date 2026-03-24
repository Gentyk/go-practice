package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func sender(url string, ctx context.Context, chan_in chan<- string) error {
	_, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	chan1 := make(chan string)

	for _, i := range urls {
		go func(url string, ctx context.Context, chan_in chan<- string) {
			sender(url, ctx, chan1)
		}(i, ctx, chan1)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-chan1)
	}
	close(chan1)
}
