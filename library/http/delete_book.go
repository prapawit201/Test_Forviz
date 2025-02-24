package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/utility"
)

// DeleteBook implements ServerInterface.
func (h *Handler) DeleteBook(c *fiber.Ctx, id string) error {
	request := entity.DeleteBookRequest{ID: id}

	//Validate input
	if err := utility.Validate(c, request); err != nil {
		log.Printf("utility.Validate DeleteBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	// call service DeleteBook
	if err := h.LibraryService.DeleteBook(c, &entity.DeleteBookRequest{ID: id}); err != nil {
		log.Printf("h.UserService.CreateBook Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusInternalServerError, err)
	}

	return utility.Response(c, http.StatusOK, nil)
}
