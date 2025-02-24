package http

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/utility"
)

// GetBookDetail implements ServerInterface.
func (h *Handler) GetBookDetail(c *fiber.Ctx, id string) error {
	request := entity.GetBookDetailRequest{ID: id}

	//Validate input
	if err := utility.Validate(c, request); err != nil {
		log.Printf("utility.Validate GetBookDetail Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	// call service GetBookDetail
	resp, err := h.LibraryService.GetBookDetail(c, &entity.GetBookDetailRequest{ID: id})
	if err != nil {
		log.Printf("h.UserService.GetBookDetail Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusInternalServerError, err)

	}

	return utility.Response(c, http.StatusOK, nil, resp)
}
