package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		select {
		case <-req.Context().Done():
			log.Println("request canceled")
			res.WriteHeader(http.StatusRequestTimeout)
			return

		case <-time.After(5 * time.Second):
			log.Println("timeout 5 seconds")
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte("hello world"))
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
