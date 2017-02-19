package common

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {

	if session == nil {

		var err error
		session, err = mgo.Dial(AppConfig.MongoDBHost)
		
		// Check if connection error, is mongo running?
		if err != nil {
			panic(err)
		}
	}

	return session
}