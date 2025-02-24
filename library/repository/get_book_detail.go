package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// GetBookDetail implements libraryRepository.
func (u *libraryRepository) GetBookDetail(ctx *fiber.Ctx, input *entity.Book) (*entity.Book, error) {
	book := model.Book{}

	tx := u.DB.Where("uuid = ?", input.UUID).Last(&book)
	if tx.Error != nil {
		log.Printf("Repositorty GetBookDetail error: %s", tx.Error.Error())
		return nil, tx.Error
	}

	return book.ToEntity(), nil
}
