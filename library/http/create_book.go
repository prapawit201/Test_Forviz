package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/utility"
)

// CreateBook implements ServerInterface.
func (h *Handler) CreateBook(c *fiber.Ctx) error {
	request := entity.CreateBookRequest{}

	if err := c.BodyParser(&request); err != nil {
		log.Printf("Context BodyParser Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	//Validate input
	if err := utility.Validate(c, request); err != nil {
		log.Printf("utility.Validate CreateBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	// call service CreateBook
	if err := h.LibraryService.CreateBook(c, &request); err != nil {
		log.Printf("h.UserService.CreateBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusInternalServerError, err)
	}

	return utility.Response(c, http.StatusCreated, nil)
}
