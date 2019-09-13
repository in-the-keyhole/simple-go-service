package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	http.HandleFunc("/api/v1/hello", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer - say Hello to name parameter
func HelloServer(w http.ResponseWriter, r *http.Request) {
	nameParameter := r.URL.Query().Get("name")
	log.WithFields(log.Fields{
		"name": nameParameter,
	}).Info("recieved /api/v1/hello")
	response := HelloResponse{Greeting: fmt.Sprintf("Hello, %s from Go", nameParameter)}
	json.NewEncoder(w).Encode(response)

}
