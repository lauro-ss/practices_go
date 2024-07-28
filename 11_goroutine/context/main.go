package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	makeRequest(ctx, "https://www.google.com.br/")

	select {
	case <-ctx.Done():
		if ctx.Err() != nil {
			if errors.Is(ctx.Err(), context.Canceled) {
				os.Stdout.WriteString("Context foi cancelado \n")
			} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				os.Stdout.WriteString("Context timeout \n")
			}
		} else {
			os.Stdout.WriteString("Context finalizou com sucesso \n")
		}
	}

	// r, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com.br/", nil)
	// if err != nil {
	// 	log.Fatal(err, errors.Is(err, http.ErrNoLocation))
	// }
	// client := http.DefaultClient
	// res, err := client.Do(r)
	// if err != nil {
	// 	log.Fatal(err, errors.Is(err, context.DeadlineExceeded))
	// }
	// fmt.Println(res.StatusCode)
}

func makeRequest(ctx context.Context, url string) {
	fmt.Println("request to", url)
	//time.Sleep(time.Second * 1)
	if ctx.Err() == nil {
		fmt.Println("response from", url)
	}
}
