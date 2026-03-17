package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))

	req.Body.Close()
}
