package api

import (
	"encoding/json"
	"fmt"

	"github.com/dimfeld/httptreemux"
	"github.com/fabianocostaalvarenga/filmes/db"
	"github.com/fabianocostaalvarenga/filmes/filmes"
	"net/http"
)

type FilmePostHandler struct{}

func (h *FilmePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	filme := &filmes.Filme{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(filme)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository := db.NewFilmeRepository()

	repository.Create(filme)

}

type FilmePutHandler struct{}

func (h *FilmePutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Alterar PUT")
}

type FilmeDeleteHandler struct{}

func (d *FilmeDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	repository := db.NewFilmeRepository()

	err := repository.Remove(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

}

type FilmeGetOneHandler struct{}

func (d *FilmeGetOneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	repository := db.NewFilmeRepository()

	filme, err := repository.FindById(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(filme)
	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao encodar o Json...")
		return
	}

}

type FilmeGetAllHandler struct{}

func (d *FilmeGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	repository := db.NewFilmeRepository()

	filmes, repositoryErr := repository.FindAll()

	if repositoryErr != nil {
		http.Error(w, repositoryErr.Error(), http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao recuperar todos os registros...")
		return
	}

	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(filmes)
	if encodeErr != nil {
		http.Error(w, encodeErr.Error(), http.StatusInternalServerError)
		fmt.Fprintln(w, "Erro ao encodar Json...")
		return
	}

}
