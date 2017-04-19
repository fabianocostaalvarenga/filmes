package main

import (
	//"errors"
	"log"

	"github.com/dimfeld/httptreemux"
	"github.com/fabianocostaalvarenga/filmes/api"
	"net/http"
)

func main() {

	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodPost, "/cadastra", &api.FilmePostHandler{})
	router.Handler(http.MethodPut, "/altera", &api.FilmePutHandler{})
	router.Handler(http.MethodDelete, "/deleta/:id", &api.FilmeDeleteHandler{})
	router.Handler(http.MethodGet, "/obtem/:id", &api.FilmeGetOneHandler{})
	router.Handler(http.MethodGet, "/lista", &api.FilmeGetAllHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

}
