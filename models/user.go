package models

import "gopkg.in/mgo.v2/bson"

type (
	// User structure
	User struct {
		Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name string        `json:"name" bson:"name"`
		Gravatar string    `json:"gravatar" bson:"gravatar"`
	}
)
