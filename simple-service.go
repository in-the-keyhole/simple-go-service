package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/hello", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer - say Hello to name parameter
func HelloServer(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	fmt.Fprintf(w, "Hello, %s!", queryValues.Get("name"))
}
