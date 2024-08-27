package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Auction represents an auction entity
type Auction struct {
	ItemID   primitive.ObjectID `bson:"itemId" json:"itemId"`
	SellerID primitive.ObjectID `bson:"sellerId" json:"sellerId"`
	EndTime  time.Time          `bson:"endTime" json:"endTime"`
	Bids     []Bid              `bson:"bids" json:"bids"`
	Status   string             `bson:"status" json:"status"`
}
