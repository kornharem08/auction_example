package auction

import (
	"context"

	"github.com/kornharem08/auction_example/lib/mongo"
	"github.com/kornharem08/auction_example/models"
)

type IRepository interface {
	Create(ctx context.Context, data models.Auction) error
}

type Repository struct {
	Collection mongo.ICollection
}

func NewRepository(dbconn mongo.IConnect) IRepository {
	return &Repository{
		Collection: mongo.NewCollection(dbconn.Database().Collection("auctions")),
	}
}

func (repo Repository) Create(ctx context.Context, data models.Auction) error {
	if _, err := repo.Collection.InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}
