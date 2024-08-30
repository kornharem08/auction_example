package config

type Config struct {
	// MongoDB
	URI                string `env:"MONGO_URI" default:"mongodb://localhost:27017/"`
	MongoDBDatabase    string `env:"MONGODB_DATABASE_NAME" default:"auction_management"`
	AuctionsCollection string `env:"MONGODB_AUCTIONS_COLLECTION_NAME" default:"auctions"`
}
