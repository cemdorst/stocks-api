package dao

import (
	"log"
	//"fmt"

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
	STOCK_COLLECTION = "all_stocks"
	HIST_COLLECTION = "historical"
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
	err := db.C(STOCK_COLLECTION).Find(bson.M{}).All(&stocks)
	return stocks, err
}

// Find historical data 
func(m * StocksDAO) FindHistoricalBySymbol(symbol string)([]Historical, error) {
    var data []Historical
    err := db.C(HIST_COLLECTION).Find(bson.M{"symbol": symbol}).All(&data)
    return data, err
}

// Find volatility data 
func(m * StocksDAO) FindVolatilityBySymbol(symbol string)([]Volatility, error) {
    var data []Volatility
    var sum float64
    var historical_data []Historical
    err := db.C(HIST_COLLECTION).Find(bson.M{"symbol": symbol}).All(&historical_data)
    for _, item := range historical_data {
	    sum = sum + item.Close
	    data = []Volatility{{Symbol: "BRSR3", Volty10: sum, Volty30: sum/float64(len(historical_data))},}
    }
    return data, err
}
