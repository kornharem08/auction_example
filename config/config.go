package config

type Config struct {
	// MongoDB
	URI                string `env:"MONGO_URI" default:"mongodb://root:example@0.0.0.0:27017/tcp/"`
	MongoDBDatabase    string `env:"MONGODB_DATABASE_NAME" default:"auction_management"`
	AuctionsCollection string `env:"MONGODB_AUCTIONS_COLLECTION_NAME" default:"auctions"`
}
