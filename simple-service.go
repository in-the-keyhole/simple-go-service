package main

import (
	"encoding/json"
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
	response := HelloResponse{Greeting: fmt.Sprintf("Hello, %s from Go", queryValues.Get("name"))}
	json.NewEncoder(w).Encode(response)

}
