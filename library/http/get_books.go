package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/utility"
	"go.openly.dev/pointy"
)

// GetBooks implements ServerInterface.
func (h *Handler) GetBooks(c *fiber.Ctx, params GetBooksParams) error {
	request := entity.GetBooksRequest{
		Page:        pointy.IntValue(params.Page, 1),   // if null set default page = 1
		Limit:       pointy.IntValue(params.Limit, 10), // if null set default size = 10
		FilterBy:    params.FilterBy,
		FilterValue: params.FilterValue,
		Sort:        params.Sort,
	}

	//Validate input
	if err := utility.Validate(c, request); err != nil {
		log.Printf("utility.Validate GetBooks Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	//validate filter
	if (request.FilterBy != nil && request.FilterValue == nil) || (request.FilterBy == nil && request.FilterValue != nil) {
		err := fmt.Errorf("%s", "Value of filter_by or filter_value must be both")
		log.Printf("utility.Validate GetBooks Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusBadRequest, err)
	}

	// call service GetBooks
	resp, err := h.LibraryService.GetBooks(c, &request)
	if err != nil {
		log.Printf("h.UserService.GetBooks Error Occurred: %s", err.Error())
		return utility.Response(c, http.StatusInternalServerError, err)

	}

	return utility.Response(c, http.StatusOK, nil, resp)
}
