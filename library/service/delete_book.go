package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
)

// DeleteBook implements libraryService.
func (u *libraryService) DeleteBook(ctx *fiber.Ctx, input *entity.DeleteBookRequest) error {
	//initial
	log.Println("Initial delete book service")

	if err := u.LibraryRepository.DeleteBook(ctx, &entity.Book{UUID: input.ID}); err != nil {
		log.Printf("LibraryRepository.DeleteBook err: %s", err.Error())
		return err
	}

	return nil
}
