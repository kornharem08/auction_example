package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kornharem08/auction_example/config"
	_ "github.com/kornharem08/auction_example/docs"
	"github.com/kornharem08/auction_example/handlers"
	"github.com/kornharem08/auction_example/lib/environ"
	"github.com/kornharem08/auction_example/lib/mong"
)

// func init() {
// 	hostname, err := os.Hostname()
// 	if err != nil {
// 		log.Fatalf("Error when get hostmane: %s", err)
// 	}

// 	log.Fatalf("Action start app  %s", hostname)
// }

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
	if cfg.MongoDBDatabase == "" {
		log.Fatal("database name must be present")
	}

	dbconn, err := mong.New(cfg.MongoDBDatabase)
	// New database connection
	if err != nil {
		log.Fatal(err)
	}

	// Ensure connection close
	defer dbconn.Close()

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
	app.Get("/auctions", handlerAuctions.GetListAuction)

	app.Listen(":3000")

}
