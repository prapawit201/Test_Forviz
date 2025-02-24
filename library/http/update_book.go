package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/utility"
)

// UpdateBook implements ServerInterface.
func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	request := entity.UpdateBookRequest{}

	if err := c.BodyParser(&request); err != nil {
		log.Printf("Context BodyParser Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	//Validate input
	if err := utility.Validate(c, request); err != nil {
		log.Printf("utility.Validate UpdateBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	// call service UpdateBook
	if err := h.LibraryService.UpdateBook(c, &request); err != nil {
		log.Printf("h.UserService.UpdateBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusInternalServerError, err)
	}

	return utility.Response(c, http.StatusOK, nil)
}
