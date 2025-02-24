package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// CreateBook implements libraryRepository.
func (u *libraryRepository) CreateBook(ctx *fiber.Ctx, input *entity.Book) error {
	book := new(model.Book).ToModelCreate(input)

	tx := u.DB.Create(&book)
	if tx.Error != nil {
		log.Printf("Repositorty CreateBook error: %s", tx.Error.Error())
		return tx.Error
	}

	return nil
}
