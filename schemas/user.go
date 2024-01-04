package schemas

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Mail     string        `bson:"mail" json:"mail"`
	Password string        `bson:"password" json:"password"`
}
