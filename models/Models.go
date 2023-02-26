package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoList struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title" json:"title"`
	Completed string             `bson:"completed" json:"completed"`
	CreatedAt time.Time          `bson:"createAt" json:"created_at"`
	Todo_list []string           `json:"todo_list"`
}
