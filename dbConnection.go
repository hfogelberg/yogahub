package main

import (
	mgo "gopkg.in/mgo.v2"
)

const MongoDb = "yogahub"

type Connection struct {
	Db *mgo.Database
}
