package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	client := http.Client{}

	r, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(r)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {

	}
	fmt.Fprintln(os.Stdout, string(body))

}
