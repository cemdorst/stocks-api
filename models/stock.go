package models

import "gopkg.in/mgo.v2/bson"

// Represents a stock, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Stock struct {
	Symbol      string        `bson:"symbol" json:"symbol"`
	Sector      string        `bson:"sector" json:"sector"`
	Subsector   string        `bson:"subSector" json:"subsector"`
}

type Historical struct {
	Symbol      string        `bson:"symbol" json:"symbol"`
	Close       float64       `json:"close"`
	Date        string        `bson:"date" json:"date"`
}

type Volatility struct {
	Symbol      string        `bson:"symbol" json:"symbol"`
	Volty10     float64       `json:"10day"`
	Volty30     float64       `json:"30day"`
}
