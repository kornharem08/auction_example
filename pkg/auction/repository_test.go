//go:build integration
// +build integration

package auction_test

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kornharem08/auction_example/lib/mong"
	mongmocks "github.com/kornharem08/auction_example/lib/mong/mocks"
	"github.com/kornharem08/auction_example/lib/mong/mongstub"
	"github.com/kornharem08/auction_example/models"
	auction "github.com/kornharem08/auction_example/pkg/auction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	collection mong.ICollection
	repo       auction.IRepository
)

const (
	COLLECTION_NAME = "auctions"
	ENV_NAME        = "MONGODB_AUCTIONS_COLLECTION_NAME"
)

func TestMain(m *testing.M) {
	container, err := mongstub.Connect("demos")
	if err != nil {
		panic(err)
	}

	defer func(ctx context.Context) {
		container.Client.Close()
		container.Terminate(ctx)

	}(context.Background())

	// Keep collection
	collection = container.Client.Database().Collection(COLLECTION_NAME)

	// New repository
	repo = auction.NewRepository(container.Client)

	os.Exit(m.Run())
}
func TestNewRepository(t *testing.T) {
	var (
		dbconn     *mongmocks.IConnect
		database   *mongmocks.IDatabase
		collection *mongmocks.ICollection
	)

	beforeEach := func() {
		os.Setenv(ENV_NAME, COLLECTION_NAME)

		dbconn = new(mongmocks.IConnect)
		database = new(mongmocks.IDatabase)
		collection = new(mongmocks.ICollection)

		// Mocking
		dbconn.On("Database").Return(
			func() mong.IDatabase {
				return database
			},
		)
		database.On("Collection", mock.Anything).Return(
			func(string) mong.ICollection {
				return collection
			},
		)
	}

	afterEach := func() {
		os.Unsetenv(ENV_NAME)
	}

	t.Run("Should be new repository", func(t *testing.T) {
		beforeEach()
		defer afterEach()

		repo := auction.NewRepository(dbconn)
		v := reflect.Indirect(reflect.ValueOf(repo))

		for index := 0; index < v.NumField(); index++ {
			assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
		}

		database.AssertCalled(t, "Collection", COLLECTION_NAME)
	})
}

func TestRepository_Create(t *testing.T) {
	var (
		ctx context.Context
	)

	beforeEach := func() {
		ctx = context.Background()
	}
	afterEach := func() {
		collection.Drop(ctx)
	}

	mockAuction := models.Auction{
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

	t.Run("Should create auction", func(t *testing.T) {
		beforeEach()
		defer afterEach()
		err := repo.Create(ctx, mockAuction)
		assert.NoError(t, err)
	})

	t.Run("Should return error when insert", func(t *testing.T) {
		beforeEach()
		defer afterEach()

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		err := repo.Create(ctx, mockAuction)
		assert.ErrorIs(t, err, context.Canceled)
	})
}
