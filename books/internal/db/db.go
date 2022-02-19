package db

import (
	"github.com/supperdoggy/grpc_CRUD/book_proto"
	"github.com/u2takey/go-utils/rand"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"
)

type IDB interface {
}

type db struct {
	l *zap.Logger

	dbname                string
	booksCollection       *mgo.Collection
	booksDataCollection   *mgo.Collection
	scoresCollection      *mgo.Collection
	currentPageCollection *mgo.Collection
}

func NewDB(l *zap.Logger, url, dbName, books, booksdata, scores, currpage string) (IDB, error) {
	client, err := mgo.Dial("")
	if err != nil {
		l.Error("error dialing with db", zap.Error(err))
		return nil, err
	}

	database := client.DB(dbName)
	result := db{
		l:                     l,
		dbname:                dbName,
		booksCollection:       database.C(books),
		booksDataCollection:   database.C(booksdata),
		scoresCollection:      database.C(scores),
		currentPageCollection: database.C(currpage),
	}

	return result, nil
}

func (d *db) NewBook(name, author string, data []byte) error {
	rand.Seed(time.Now().UnixNano())
	book := book_proto.Book{
		Id: rand.String(24),
		Name: name,
		Author: author,
		CreationDate: time.Now().Unix(),
		UploadTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}
