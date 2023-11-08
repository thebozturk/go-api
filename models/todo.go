package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Todo is a struct for todo
// json:"title" is a tag for json for swagger etc.
// primitive.ObjectID is a type for mongodb
type Todo struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title.omitempty"`
	Content string             `json:"content,omitempty"`
}
