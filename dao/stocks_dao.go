package dao

import (
	"log"

	. "github.com/cemdorst/stocks-api/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type StocksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "all_stocks"
)

// Establish a connection to database
func (m *StocksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of stocks
func (m *StocksDAO) FindAll() ([]Stock, error) {
	var stocks []Stock
	err := db.C(COLLECTION).Find(bson.M{}).All(&stocks)
	return stocks, err
}
