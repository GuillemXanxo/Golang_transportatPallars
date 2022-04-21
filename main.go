package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getViatgesCrono(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(res, "You are trying to get the trips")
}

func main() {
  port := ":4040"
  router := mux.NewRouter().StrictSlash(true)
  

  router.HandleFunc("/viatges/crono", getViatgesCrono)
  log.Fatal(http.ListenAndServe(port, router))
}
