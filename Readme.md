## Docker Run Mongo
    1. run docker-compose up -d

## Run Service 
    1. make run

## Swagger init
    1. go install github.com/swaggo/swag/cmd/swag@latest
    1. run swag init -g cmd/main.go


## Data 
    {
    "id": "64e04c12a3f5b7b6d50a4e42",
    "itemId": "64e04c12a3f5b7b6d50a4e31",
    "sellerId": "64e04c12a3f5b7b6d50a4e32",
    "endTime": "2024-08-26T13:00:00Z",
    "bids": [
        {
            "userId": "64e04c12a3f5b7b6d50a4e33",
            "bidAmount": 100.0,
            "bidTime": "2024-08-26T12:00:00Z"
        },
        {
            "userId": "64e04c12a3f5b7b6d50a4e34",
            "bidAmount": 150.0,
            "bidTime": "2024-08-26T12:15:00Z"
        }
    ],
    "status": "active"
    }