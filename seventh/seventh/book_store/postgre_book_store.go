package book_store

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"strconv"
)

type PostgreConfig struct {
	User string
	Password string
	Port string
	Host string
}

type postgreStore struct {
	db *pg.DB
}

func NewPostgreBookStore(config PostgreConfig) (BookStore, error) {
	db := pg.Connect(&pg.Options{
		Addr: config.Host + ":" + config.Port,
		User: "postgres",
		Password: config.Password,
	})
	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &postgreStore{db: db}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Book)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *postgreStore) Create(book *Book) (*Book, error) {
	return book, ps.db.Insert(book)
}

func (ps *postgreStore) GetBook()

