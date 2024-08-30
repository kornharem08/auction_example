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
	"go.mongodb.org/mongo-driver/bson"
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

func TestRepository_GetList(t *testing.T) {
	var (
		ctx context.Context
	)

	beforeEach := func() {
		ctx = context.Background()
		groups := []any{
			models.Auction{
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
						BidTime:   time.Now().Add(15 * time.Minute),
					},
				},
				Status: "active",
			},
			models.Auction{
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
						BidTime:   time.Now().Add(15 * time.Minute),
					},
				},
				Status: "active",
			},
		}

		_, err := collection.InsertMany(ctx, groups)
		assert.NoError(t, err)
	}

	afterEach := func() {
		collection.Drop(ctx)
	}

	t.Run("Should retrieve all auctions", func(t *testing.T) {
		beforeEach()
		defer afterEach()

		auctions, err := repo.GetList(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(auctions))
	})

	t.Run("Should return empty list when no auctions found", func(t *testing.T) {
		// No beforeEach call here to keep the collection empty
		afterEach()

		// Ensure the collection is empty
		count, err := collection.CountDocuments(ctx, bson.M{})
		assert.NoError(t, err)
		assert.Equal(t, int64(0), count)

		// Call GetList and expect an empty list
		auctions, err := repo.GetList(ctx)
		assert.NoError(t, err)
		assert.Empty(t, auctions)
	})

	t.Run("Should return error when get list", func(t *testing.T) {
		beforeEach()
		defer afterEach()

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := repo.GetList(ctx)
		assert.ErrorIs(t, err, context.Canceled)
	})

}
