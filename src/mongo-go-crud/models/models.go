package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Body      string             `json:"title" bson:"title,omitempty"`
}
