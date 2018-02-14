package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Pose struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	TitleEn     string        `json:"title-en" bson:"titleEn"`
	TitleSa     string        `json:"title-sa" bson:"titleSa"`
	Description string        `json:"description" bson:"description"`
	File        string        `json:"file" bson:"file"`
	When        time.Time     `json:"when" bson:"when"`
}
