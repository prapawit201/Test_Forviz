package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"gorm.io/gorm"
)

//go:generate /Users/babe/go/bin/mockgen -source=init.go -destination=mocks/repository.go

type LibraryRepository interface {
	//book
	CreateBook(ctx *fiber.Ctx, input *entity.Book) error
	UpdateBook(ctx *fiber.Ctx, input *entity.Book) error
	UpdateBookStatus(ctx *fiber.Ctx, input *entity.Book) error
	GetBooks(ctx *fiber.Ctx, input *entity.Book, paginate ...entity.BookPaginate) (int, []entity.Book, error)
	GetBookDetail(ctx *fiber.Ctx, input *entity.Book) (*entity.Book, error)
	DeleteBook(ctx *fiber.Ctx, input *entity.Book) error

	//book borrow log
	CreateBookBorrowLog(ctx *fiber.Ctx, input *entity.BookBorrowLog) error
}

type libraryRepository struct {
	DB *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) LibraryRepository {
	return &libraryRepository{
		DB: db,
	}
}
