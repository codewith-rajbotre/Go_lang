package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

//Notes :
// ✅ json:"_id,omitempty" ensures JSON serialization compatibility.
// ✅ bson:"_id,omitempty" ensures MongoDB stores the field as _id.
// ✅ omitempty prevents empty values from appearing in JSON/BSON output.
// ✅ MongoDB automatically generates an _id if none is provided.
