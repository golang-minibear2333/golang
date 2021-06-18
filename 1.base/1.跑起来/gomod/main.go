package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
	log.WithFields(log.Fields{
		"asd": "asd",
	}).Info("qewr")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
