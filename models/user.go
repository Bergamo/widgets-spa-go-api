package models

import "gopkg.in/mgo.v2/bson"

type (
	// User structure
	User struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        `json:"name" bson:"name"`
	}
)
