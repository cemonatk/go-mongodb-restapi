package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Id      uint8  `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name" binding:"required"`
	Address string `json:"address" bson:"address" binding:"required"`
	Age     int    `json:"age" bson:"age" binding:"required"`
	Gender  string `json:"gender" bson:"gender"`
}
