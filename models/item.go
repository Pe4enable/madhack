package models

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"time"
)

type Item struct {
	ID       objectid.ObjectID `json:"_id" bson:"_id,omitempty"`
	UDID     string            `json:"udid" bson:"udid,omitempty"`
	Email    string            `json:"email"`
	Name     string            `json:"name"`
	Surname  string            `json:"surname"`
	Created  time.Time         `json:"-"`
}
