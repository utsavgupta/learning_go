package main

import (
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("<h1>I am a little awesome cool website</h1>"))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
