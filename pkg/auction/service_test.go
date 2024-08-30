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

func TestGetList_Service(t *testing.T) {
	var (
		mockRepo  *mocks.IRepository
		service   auction.IService
		ctx       context.Context
		errCreate error
		data      []models.Auction
	)

	beforeEach := func() {
		mockRepo = new(mocks.IRepository)
		service = auction.Service{
			Repository: mockRepo,
		}

		data = []models.Auction{
			{
				ItemID:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
				SellerID: [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
				EndTime:  time.Now(),
				Bids: []models.Bid{
					{
						UserID:    [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
						BidAmount: 100.0,
						BidTime:   time.Now(),
					},
					{
						UserID:    [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
						BidAmount: 150.0,
						BidTime:   time.Now().Add(15 * time.Minute), // Adds 15 minutes to the current time
					},
				},
				Status: "active", // Example status
			},
		}

		// Mocking
		mockRepo.On("GetList", mock.Anything).Return(
			func(context.Context) []models.Auction {
				return data
			},
			func(context.Context) error {
				return errCreate
			},
		)

	}

	t.Run("Should return auctions list", func(t *testing.T) {
		beforeEach()
		// Mock data for Auctions list
		expectedResult := []models.Auction{
			{
				ItemID:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
				SellerID: [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
				EndTime:  time.Now(),
				Bids: []models.Bid{
					{
						UserID:    [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
						BidAmount: 100.0,
						BidTime:   time.Now(),
					},
					{
						UserID:    [12]uint8{0, 0, 0, 0, 0, 0, 0, 1},
						BidAmount: 150.0,
						BidTime:   time.Now().Add(15 * time.Minute), // Adds 15 minutes to the current time
					},
				},
				Status: "active", // Example status
			},
		}

		list, err := service.GetList(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, list)
	})

	t.Run("Should error when GetList", func(t *testing.T) {
		beforeEach()
		data = []models.Auction{}
		errCreate = errors.New("data not found")

		list, err := service.GetList(ctx)
		assert.ErrorIs(t, err, errCreate)

		assert.Empty(t, list)
	})
}
