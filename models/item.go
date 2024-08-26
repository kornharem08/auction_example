package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Item represents an item that is being auctioned
type Item struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Description   string             `bson:"description"`
	StartingPrice float64            `bson:"startingPrice"`
	CurrentPrice  float64            `bson:"currentPrice"`
	Images        []string           `bson:"images"`
}
