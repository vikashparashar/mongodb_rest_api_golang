package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id , omitempty" json:"id , omitempty"`
	Name string `bson:"name , omitempty" json:"name , omitempty"`
	Gender string `bson:"gender , omitempty" json:"gender , omitempty"`
	Age int64 `bson:"age , omitempty" json:"age , omitempty"`
}