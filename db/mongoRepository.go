package db

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"github.com/fabianocostaalvarenga/filmes/filmes"
)

type Filme struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Ano  int64  `bson:ano`
}

const FilmeCollection = "filme"

var ErrFilmes = errors.New("Erro na manipulação de filme!")

type FilmeRepository struct {
	session *mgo.Session
}

func (r *FilmeRepository) Create(p *Filme) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(FilmeCollection)
	err := collection.Insert(p)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrFilmes
		}
	}
	return err
}

func (r *FilmeRepository) Update(p *Filme) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(FilmeCollection)
	return collection.Update(bson.M{"_id": p.Id}, p)
}

func (r *FilmeRepository) Remove(id string) error {

	log.Println("Removendo usuario " + id)

	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(FilmeCollection)

	query := bson.M{"_id": id}
	err := collection.Remove(query)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Usuario " + id + " removido com sucesso!")
	}

	return err
}

func (r *FilmeRepository) FindAll() ([]*Filme, error) {

	log.Println("Recuperando todos os usuários...")

	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(FilmeCollection)

	documents := make([]*Filme, 0)

	err := collection.Find(nil).All(&documents)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(len(documents), " Usuário(s) retornado(s) com sucesso!")
	}

	return documents, err
}

func (r *FilmeRepository) FindById(id string) (*Filme, error) {

	log.Println("Buscando usuario " + id)

	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(FilmeCollection)
	query := bson.M{"_id": id}

	filme := &Filme{}

	err := collection.Find(query).One(filme)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Usuario " + id + " encontrado...")
	}

	return filme, err
}

func NewFilmeRepository() *FilmeRepository {
	session, err := mgo.Dial("localhost:27017/go-course-filme")

	if err != nil {
		log.Fatal(err)
	}

	return &FilmeRepository{session}
}
