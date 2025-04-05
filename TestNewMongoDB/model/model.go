package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"movie,omitempty"`
	Department string             `json:"department,omitempty"`
}
