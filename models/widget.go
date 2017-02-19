package models

import "gopkg.in/mgo.v2/bson"

type (
	// User structure
	Widget struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        `json:"name" bson:"name"`
		Color string       `json:"color" bson:"color"`
        Price string       `json:"price" bson:"price"`
        Inventory int      `json:"inventory" bson:"inventory"`
        Melts bool         `json:"melts" bson:"melts"`
	}
)