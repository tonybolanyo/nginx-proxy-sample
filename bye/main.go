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
	router.HandleFunc("/", ByeByeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func ByeByeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, bye, I'm another golang microservice")
	w.WriteHeader(http.StatusAccepted) //RETURN HTTP CODE 202
}
