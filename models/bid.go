package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bid represents a bid made by a user
type Bid struct {
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	BidAmount float64            `bson:"bidAmount" json:"bidAmount"`
	BidTime   time.Time          `bson:"bidTime" json:"bidTime"`
}
