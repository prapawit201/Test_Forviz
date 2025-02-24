package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
)

// UpdateBookStatus implements LibraryRepository.
func (u *libraryRepository) UpdateBookStatus(ctx *fiber.Ctx, input *entity.Book) error {

	tx := u.DB.Model(&model.Book{}).Where("uuid = ?", input.UUID)

	if input.IsBorrow {
		tx = tx.Updates(map[string]interface{}{"is_borrow": input.IsBorrow, "borrow_count": input.BorrowCount})
	} else {
		tx = tx.Update("is_borrow", input.IsBorrow)
	}

	if tx.Error != nil {
		log.Printf("Repositorty UpdateBookStatus error: %s", tx.Error.Error())
		return tx.Error
	}

	return nil
}
