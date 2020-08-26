package models

import "gopkg.in/mgo.v2/bson"

// Represents a stock, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Stock struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Symbol      string        `bson:"symbol" json:"symbol"`
	Sector      string        `bson:"sector" json:"sector"`
	Subsector   string        `bson:"subsector" json:"subsector"`
}
