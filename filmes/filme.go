package filmes

type Filme struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
	Ano  int64  `bson:ano`
}
