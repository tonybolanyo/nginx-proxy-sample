//main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HelloHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I'm a golang microservice")
	w.WriteHeader(http.StatusAccepted) //RETURN HTTP CODE 202
}
