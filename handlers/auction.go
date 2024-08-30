package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kornharem08/auction_example/lib/mong"
	"github.com/kornharem08/auction_example/models"
	"github.com/kornharem08/auction_example/pkg/auction"
)

type IHandler interface {
	CreateAuction(c *fiber.Ctx) error
	GetListAuction(c *fiber.Ctx) error
}

type Handler struct {
	auctionsService auction.IService
}

func NewHandler(dbconn mong.IConnect) IHandler {
	return &Handler{
		auctionsService: auction.NewService(dbconn),
	}
}

// CreateAuction creates a new auction
// @Summary Create a new auction
// @Description Create a new auction in the system
// @Tags auctions
// @Accept json
// @Produce json
// @Param auction body models.Auction true "Auction details"
// @Success 201 {object} models.Auction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auctions [post]
func (h *Handler) CreateAuction(c *fiber.Ctx) error {
	var newAuction models.Auction

	// Log incoming request body for debugging
	body := c.Body()
	log.Printf("Received body: %s", body)

	// Parse the JSON body into the Auction struct
	if err := c.BodyParser(&newAuction); err != nil {
		log.Printf("Error parsing JSON: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Log parsed auction data
	log.Printf("Parsed Auction: %+v", newAuction)

	// Create the auction
	err := h.auctionsService.Create(c.Context(), newAuction)
	if err != nil {
		log.Printf("Error creating auction: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create auction"})
	}

	return c.Status(fiber.StatusCreated).JSON(newAuction)
}

// GetListAuction godoc
// @Summary Get list of auctions
// @Description Retrieves a list of all auctions
// @Tags Auctions
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Auction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auctions [get]
func (h *Handler) GetListAuction(c *fiber.Ctx) error {
	// Create a context
	ctx := c.Context()

	// Call the GetList method from the service
	auctions, err := h.auctionsService.GetList(ctx)
	if err != nil {
		// Return a 500 Internal Server Error if something goes wrong
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve auctions",
		})
	}

	// Return the list of auctions as a JSON response
	return c.JSON(auctions)
}
