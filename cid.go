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
	_"encoding/json"
	"fmt"
)

const SERVICE_1 = `{
		"name": "Service1",
		"definition": "some definition stuff",
		"id": 1
	}`

func stubbedHandler(requestDescription string, responseText string) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("[REQUEST] " + requestDescription)

		fmt.Printf("[RESPONSE] %s", responseText)
		fmt.Fprintf(w, responseText)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/services", stubbedHandler("All Services", "[" + SERVICE_1 + "]"))
	http.HandleFunc("/service/1", stubbedHandler("Service 1", SERVICE_1))

	fmt.Println("Listening on port 5678")
	http.ListenAndServe(":5678", nil)
}
