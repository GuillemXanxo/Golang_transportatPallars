package main

import (
	"github.com/gorilla/mux"
)

type Viatge struct {
  origen string `json:origen`
  desti string `json:desti`
  places int `json:places`
  dataNumber int
  data string
  horaSortidaNumber int
  horaSortida string
  comentaris string
  dones bool
}

func main() {
  mux.NewRouter()
}
