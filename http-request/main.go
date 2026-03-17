package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	client := http.Client{}

	requist, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	requist.Header.Set("Accept", "application/json")

	resp, err := client.Do(requist)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
	fmt.Print(string(bod))
}
