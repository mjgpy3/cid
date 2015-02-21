//> # `cid`
//> Continuous Integration in Docker

//> Project was built using docker versions:
//>  - Client version: 1.5.0
//>  - Client API version: 1.17
//>  - Go version (client): go1.4.1
//>  - Git commit (client): a8a31ef

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
