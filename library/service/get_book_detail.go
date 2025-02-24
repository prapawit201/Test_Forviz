package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
)

// GetBookDetail implements libraryService.
func (u *libraryService) GetBookDetail(ctx *fiber.Ctx, input *entity.GetBookDetailRequest) (*entity.GetBookDetailResponse, error) {
	//initial
	log.Println("Initial get book detail service")

	resp, err := u.LibraryRepository.GetBookDetail(ctx, &entity.Book{UUID: input.ID})
	if err != nil {
		log.Printf("LibraryRepository.GetBookDetail err: %s", err.Error())
		return nil, err
	}

	return new(entity.GetBookDetailResponse).ToEntity(resp), nil
}
