package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IConnect interface {
	Database() *mongo.Database
	Close(ctx context.Context) error
}

type Connect struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewConnection(uri, dbName string) IConnect {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	return &Connect{
		client:   client,
		database: client.Database(dbName),
	}
}

func (c *Connect) Database() *mongo.Database {
	return c.database
}

func (c *Connect) Close(ctx context.Context) error {
	if err := c.client.Disconnect(ctx); err != nil {
		return err
	}
	log.Println("Disconnected from MongoDB")
	return nil
}
