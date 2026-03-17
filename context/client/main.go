package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	r, _ := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	select {
	case <-r.Context().Done():
		log.Println("time out expired")
		return
	case <-time.After(5 * time.Second):
		log.Println("time out 5 request ok ")
		return
	}
}
