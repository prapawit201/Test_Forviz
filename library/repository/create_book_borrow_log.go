package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// CreateBookBorrowLog implements LibraryRepository.
func (u *libraryRepository) CreateBookBorrowLog(ctx *fiber.Ctx, input *entity.BookBorrowLog) error {
	bookBorrowLog := new(model.BookBorrowLog).ToModelCreate(input)

	tx := u.DB.Create(&bookBorrowLog)
	if tx.Error != nil {
		log.Printf("Repositorty CreateBookBorrowLog error: %s", tx.Error.Error())
		return tx.Error
	}

	return nil
}
