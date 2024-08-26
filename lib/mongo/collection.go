package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
}

type Collection struct {
	collection *mongo.Collection
}

func NewCollection(col *mongo.Collection) ICollection {
	return &Collection{
		collection: col,
	}
}

func (c *Collection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(ctx, document)
}
