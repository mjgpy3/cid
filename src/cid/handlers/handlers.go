package handlers

import (
	"net/http"
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

var URL_TO_HANDLER = map[string]func(http.ResponseWriter, *http.Request) {
	"/services": stubbedHandler("All Services", "[" + SERVICE_1 + "]"),
	"/service/1": stubbedHandler("Service 1", SERVICE_1),
}
