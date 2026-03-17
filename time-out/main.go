package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}

	jsonB := bytes.NewBuffer([]byte(`{"name": "lucas"}`))
	by := []byte(`{"name": "lucas"}`)
	resp, err := c.Post("https://google.com", "application/json", jsonB)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(body))

	io.CopyBuffer(os.Stdout, resp.Body, by)
}
