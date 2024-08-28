export MONGODB_DATABASE_NAME=auction_management
export MONGODB_AUCTIONS_COLLECTION_NAME=auctions
export MONGO_URI=mongodb://root:example@0.0.0.0:27017/tcp/
test:
	go test ./... -count=1

run:
	go run cmd/main.go
