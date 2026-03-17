package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()

	ctx, cancel := context.WithTimeout(c, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("time out expired")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("time out 5")
		return
	}
}
