package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Resp struct {
	rcode int
	url   string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	urls := []string{
		"https://google.com",
		"https://www.sme.sk",
		"https://deelay.me/5000/www.nzz.ch",
	}

	results := make(chan Resp)

	for _, url := range urls {
		chkurl(ctx, url, results)
	}

	for range urls {
		resp := <-results
		fmt.Printf("Received: %d %s\n", resp.rcode, resp.url)
	}

}

func chkurl(ctx context.Context, url string, results chan Resp) {
	fmt.Printf("Fetching %s\n", url)
	httpch := make(chan int)
	go func() {

		// asynch URL fetch
		go func() {
			resp, err := http.Get(url)
			if err != nil {
				httpch <- 500
			} else {
				httpch <- resp.StatusCode
			}
		}()

		select {
		case result := <-httpch:
			results <- Resp{
				rcode: result,
				url:   url,
			}
		case <-ctx.Done():
			fmt.Printf("Timeout!\n")
			results <- Resp{
				rcode: 501,
				url:   url,
			}
		}
	}()
}
