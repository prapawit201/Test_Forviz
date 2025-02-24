package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
)

// CreateBook implements libraryService.
func (u *libraryService) CreateBook(ctx *fiber.Ctx, input *entity.CreateBookRequest) error {
	//initial
	log.Println("Initial create book service")

	if err := u.LibraryRepository.CreateBook(ctx, input.ToEntityCreate()); err != nil {
		log.Printf("LibraryRepository.CreateBook err: %s", err.Error())
		return err
	}

	return nil
}
