package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
)

// UpdateBook implements libraryService.
func (u *libraryService) UpdateBook(ctx *fiber.Ctx, input *entity.UpdateBookRequest) error {
	//initial
	log.Println("Initial UpdateBook service")

	// get book
	_, err := u.LibraryRepository.GetBookDetail(ctx, &entity.Book{UUID: input.ID})
	if err != nil {
		log.Printf("LibraryRepository.GetBookDetail err: %s", err.Error())
		return err
	}

	// update book
	if err := u.LibraryRepository.UpdateBook(ctx, input.ToEntityUpdate()); err != nil {
		log.Printf("LibraryRepository.UpdateBook err: %s", err.Error())
		return err
	}

	return nil
}
