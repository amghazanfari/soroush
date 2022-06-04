package main

import (
	"log"
	"net/http"

	"github.com/amghazanfari/soroush/src/producer"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/writemessage", producer.ProduceMessage)
	log.Fatal(http.ListenAndServe(":5065", r))
}
