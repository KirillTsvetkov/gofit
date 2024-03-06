package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Age       int                `bson:"age"`
	Weight    float32            `bson:"weight"`
	Height    int                `bson:"height"`
	DeletedAt time.Time          `json:"-" bson:"deleted_at,omitempty"`
}
