package main

import (
	"net/http"
	"fmt"
	"cid/handlers"
)


func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	for url, handler := range handlers.URL_TO_HANDLER {
		http.HandleFunc(url, handler)
	}

	fmt.Println("Listening on port 5678")
	http.ListenAndServe(":5678", nil)
}
