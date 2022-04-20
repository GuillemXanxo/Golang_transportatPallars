package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Viatge struct {
  origen string `json:origen`
  desti string `json:desti`
  places int `json:places`
  dataNumber int `json:dataNumber`
  data string `json:data`
  horaSortidaNumber int `json:horaSortidaNumber`
  horaSortida string `json:horaSortida`
  comentaris string `json:comentaris`
  dones bool `json:dones`
}

func getViatgesCrono(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "You are trying to get the trips")
}

func main() {
  port := ":4040"
  router := mux.NewRouter().StrictSlash(true)
  

  router.HandleFunc("/viatges/crono", getViatgesCrono)
  log.Fatal(http.ListenAndServe(port, router))
}
