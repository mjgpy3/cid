package main

import (
	"net/http"
	"fmt"
	"cid/handlers"
	"os"
)


func main() {
	staticDir := http.Dir(os.Getenv("GOPATH") + "/src/cid/static")
	fmt.Println("Static content from: " + staticDir)
	fs := http.FileServer(staticDir)
	http.Handle("/", fs)

	for url, handler := range handlers.URL_TO_HANDLER {
		http.HandleFunc(url, handler)
	}

	fmt.Println("Listening on port 5678")
	http.ListenAndServe(":5678", nil)
}
