export MONGODB_DATABASE_NAME=auction_management
export MONGODB_AUCTIONS_COLLECTION_NAME=auctions
export MONGO_URI=mongodb://localhost:27017/
test:
	go test ./... -count=1

run:
	go run cmd/main.go

test-integration:
	go test ./... -tags=integration -count=1
