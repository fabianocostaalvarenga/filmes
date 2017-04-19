package api

import (
	"fmt"

	"github.com/dimfeld/httptreemux"
	"github.com/fabianocostaalvarenga/filmes/db"
	"net/http"
)

type FilmePostHandler struct{}

func (h *FilmePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//params := httptreemux.ContextParams(r.Context())
	//fmt.Fprintf(w, "Eu deveria criar um carro chamado: %s!", params["id"])
	fmt.Fprintln(w, "Criar POST")
}

type FilmePutHandler struct{}

func (h *FilmePutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//params := httptreemux.ContextParams(r.Context())
	//fmt.Fprintf(w, "Eu deveria busca um carro chamado: %s!", params["id"])
	fmt.Fprintln(w, "Alterar PUT")
}

type FilmeDeleteHandler struct{}

func (d *FilmeDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	repository := db.NewFilmeRepository()

	repository.Remove(params["id"])

	fmt.Fprintln(w, "Deletar DELETE = ", params["id"])
}

type FilmeGetOneHandler struct{}

func (d *FilmeGetOneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	repository := db.NewFilmeRepository()

	repository.FindById(params["id"])

	fmt.Fprintln(w, "Obter 1 GET = ", params["id"])
}

type FilmeGetAllHandler struct{}

func (d *FilmeGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//params := httptreemux.ContextParams(r.Context())

	repository := db.NewFilmeRepository()

	repository.FindAll()

	fmt.Fprintln(w, "Obter Todos GET")
}
