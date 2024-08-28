package auction

import (
	"context"

	"github.com/kornharem08/auction_example/lib/mong"
	"github.com/kornharem08/auction_example/models"
)

type IService interface {
	Create(ctx context.Context, data models.Auction) error
}

type Service struct {
	Repository IRepository
}

func NewService(dbconn mong.IConnect) IRepository {
	return &Service{
		Repository: NewRepository(dbconn),
	}
}

func (service Service) Create(ctx context.Context, data models.Auction) error {
	return service.Repository.Create(ctx, data)
}
