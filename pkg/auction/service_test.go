package auction_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kornharem08/auction_example/models"
	"github.com/kornharem08/auction_example/pkg/auction"
	"github.com/kornharem08/auction_example/pkg/auction/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestService_Create(t *testing.T) {
	var (
		mockRepo  *mocks.IRepository
		service   auction.IService
		ctx       context.Context
		errCreate error
		data      models.Auction
	)

	beforeEach := func() {
		mockRepo = new(mocks.IRepository)
		service = auction.Service{
			Repository: mockRepo,
		}

		// Mock data for Auction
		data = models.Auction{
			ItemID:   primitive.NewObjectID(),
			SellerID: primitive.NewObjectID(),
			EndTime:  time.Now(),
			Bids: []models.Bid{
				{
					UserID:    primitive.NewObjectID(),
					BidAmount: 100.0,
					BidTime:   time.Now(),
				},
				{
					UserID:    primitive.NewObjectID(),
					BidAmount: 150.0,
					BidTime:   time.Now().Add(15 * time.Minute), // Adds 15 minutes to the current time
				},
			},
			Status: "active", // Example status
		}

		ctx = context.Background()

		// Mocking
		mockRepo.On("Create", mock.Anything, mock.Anything).Return(
			func(context.Context, models.Auction) error {
				return errCreate
			},
		)
	}

	t.Run("Should create auction", func(t *testing.T) {
		beforeEach()

		err := service.Create(ctx, data)
		assert.NoError(t, err)

		mockRepo.AssertCalled(t, "Create", ctx, data)
	})

	t.Run("Should error", func(t *testing.T) {
		beforeEach()
		errCreate = errors.New("test error creating")
		err := service.Create(ctx, data)

		assert.ErrorIs(t, err, errCreate)

	})

}
