package main

import (
	"context"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kornharem08/auction_example/config"
	_ "github.com/kornharem08/auction_example/docs"
	"github.com/kornharem08/auction_example/handlers"
	"github.com/kornharem08/auction_example/lib/mongo"
	"github.com/kornharem08/auction_example/lib/mongo/environ"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	cfg := environ.Load[config.Config]()
	dbconn := mongo.NewConnection("mongodb://localhost:27017", cfg.MongoDBDatabase)
	defer dbconn.Close(context.Background())

	app := fiber.New()
	app.Use(cors.New())

	handlerAuctions := handlers.NewHandler(dbconn)
	// Initialize Swagger middleware
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))

	// Define routes
	app.Post("/auctions", handlerAuctions.CreateAuction)

	app.Listen(":3000")

}
