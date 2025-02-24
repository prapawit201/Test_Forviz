package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// UpdateBook implements libraryRepository.
func (u *libraryRepository) UpdateBook(ctx *fiber.Ctx, input *entity.Book) error {
	book := new(model.Book).ToModelUpdate(input)

	tx := u.DB.Model(&model.Book{}).Where("uuid = ?", input.UUID).Updates(&book)
	if tx.Error != nil {
		log.Printf("Repositorty UpdateBook error: %s", tx.Error.Error())
		return tx.Error
	}

	return nil
}
