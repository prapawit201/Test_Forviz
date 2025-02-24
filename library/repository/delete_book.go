package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// DeleteBook implements libraryRepository.
func (u *libraryRepository) DeleteBook(ctx *fiber.Ctx, input *entity.Book) error {
	book := model.Book{}

	tx := u.DB.Where("uuid = ?", input.UUID).Delete(&book)

	if tx.Error != nil {
		log.Printf("Repositorty DeleteBook error: %s", tx.Error.Error())
		return tx.Error
	}

	return nil
}
