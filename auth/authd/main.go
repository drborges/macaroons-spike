package main

import (
	"github.com/drborges/macaroons-spike/auth"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	auth.Register(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}