package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetSession of mongo
func GetSession() *mgo.Session {

	if session == nil {

		var err error
		session, err = mgo.Dial(AppConfig.MongoDBHost)

		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}

	return session
}
