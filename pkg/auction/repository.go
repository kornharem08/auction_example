package auction

import (
	"context"

	"github.com/kornharem08/auction_example/config"
	"github.com/kornharem08/auction_example/lib/environ"
	"github.com/kornharem08/auction_example/lib/mong"
	"github.com/kornharem08/auction_example/models"
	"go.mongodb.org/mongo-driver/bson"
)

type IRepository interface {
	Create(ctx context.Context, data models.Auction) error
	GetList(ctx context.Context) ([]models.Auction, error)
}

type Repository struct {
	Collection mong.ICollection
}

func NewRepository(dbconn mong.IConnect) IRepository {
	return &Repository{
		Collection: dbconn.Database().Collection(environ.Load[config.Config]().AuctionsCollection),
	}
}

func (repo Repository) Create(ctx context.Context, data models.Auction) error {
	if _, err := repo.Collection.InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}

func (repo Repository) GetList(ctx context.Context) ([]models.Auction, error) {

	var (
		result []models.Auction
	)

	// Filter
	filter := bson.D{}

	// Find
	if err := repo.Collection.Find(ctx, &result, filter); err != nil {
		return result, err
	}

	// Done
	return result, nil
}
