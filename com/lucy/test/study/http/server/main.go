package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("req:ping")
		writer.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
