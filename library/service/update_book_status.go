package service

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/types/constant"
)

// UpdateBookStatus implements LibraryService.
func (u *libraryService) UpdateBookStatus(ctx *fiber.Ctx, input *entity.UpdateBookStatusRequest) error {
	//initial
	log.Println("Initial UpdateBookStatus service")

	// get book
	book, err := u.LibraryRepository.GetBookDetail(ctx, &entity.Book{UUID: input.ID})
	if err != nil {
		log.Printf("LibraryRepository.GetBookDetail err: %s", err.Error())
		return err
	}

	if (book.IsBorrow && input.Status == constant.BookStatusBorrow) || (!book.IsBorrow && input.Status == constant.BookStatusReceive) {
		err := fmt.Errorf("%s", "request status is same book current status")
		log.Printf("err: %s", err.Error())
		return err
	}

	// update status book
	if err := u.LibraryRepository.UpdateBookStatus(ctx, input.ToEntityUpdate(book)); err != nil {
		log.Printf("LibraryRepository.UpdateBook err: %s", err.Error())
		return err
	}

	// มี process borrow/receive
	if err := u.LibraryRepository.CreateBookBorrowLog(ctx, input.ToEntityCreateBookBorrowLog(book.ID)); err != nil {
		log.Printf("LibraryRepository.CreateBookBorrowLog err: %s", err.Error())
		return err
	}

	return nil
}
