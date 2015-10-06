package main

import (
	"github.com/drborges/macaroons-spike/storage"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	storage.Register(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
