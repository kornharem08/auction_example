package auction

import (
	"context"

	"github.com/kornharem08/auction_example/lib/mong"
	"github.com/kornharem08/auction_example/models"
)

type IService interface {
	Create(ctx context.Context, data models.Auction) error
	GetList(ctx context.Context) ([]models.Auction, error)
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

func (service Service) GetList(ctx context.Context) ([]models.Auction, error) {
	return service.Repository.GetList(ctx)
}
