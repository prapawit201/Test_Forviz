package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
)

// GetBooks implements libraryService.
func (u *libraryService) GetBooks(ctx *fiber.Ctx, input *entity.GetBooksRequest) (*entity.GetBooksResponse, error) {
	//initial
	log.Println("Initial get books service")

	totalCount, resp, err := u.LibraryRepository.GetBooks(ctx, input.ToEntityWhereClauseBook(), input.ToEntityPaginate())
	if err != nil {
		log.Printf("LibraryRepository.GetBooks err: %s", err.Error())
		return nil, err
	}

	// not found return empty
	if totalCount < 1 {
		log.Printf("total count book found: %d", totalCount)
		return new(entity.GetBooksResponse).ToEntityResponse(input, totalCount, nil), nil
	}

	return new(entity.GetBooksResponse).ToEntityResponse(input, totalCount, resp), nil
}
