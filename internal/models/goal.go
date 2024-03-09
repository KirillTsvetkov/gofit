package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Goal struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id"`
	Exercise Exercise           `bson:"exercise"`
	Date     time.Time          `bson:"date"`
	Status   GoalStatus         `bson:"status"`
}
