package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})
	mux.Handle("/ts", Teste{name: "lucas"})
	log.Fatal(http.ListenAndServe(":8080", mux))

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/ts", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("Hello World mux 2"))
	})

	log.Fatal(http.ListenAndServe(":8088", mux2))
}

type Teste struct {
	name string
}

func (t Teste) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello my name is " + t.name))
}
