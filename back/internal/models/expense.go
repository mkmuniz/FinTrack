package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
	Notes    string             `bson:"notes,omitempty" json:"notes,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}