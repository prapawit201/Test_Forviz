package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository"
)

//go:generate /Users/babe/go/bin/mockgen -source=init.go -destination=mocks/user.go

type LibraryService interface {
	CreateBook(ctx *fiber.Ctx, input *entity.CreateBookRequest) error
	UpdateBook(ctx *fiber.Ctx, input *entity.UpdateBookRequest) error
	UpdateBookStatus(ctx *fiber.Ctx, input *entity.UpdateBookStatusRequest) error
	GetBooks(ctx *fiber.Ctx, input *entity.GetBooksRequest) (*entity.GetBooksResponse, error)
	GetBookDetail(ctx *fiber.Ctx, input *entity.GetBookDetailRequest) (*entity.GetBookDetailResponse, error)
	DeleteBook(ctx *fiber.Ctx, input *entity.DeleteBookRequest) error
}

type libraryService struct {
	LibraryRepository repository.LibraryRepository
}

func NewLibraryService(libraryRepository repository.LibraryRepository) LibraryService {
	return &libraryService{
		LibraryRepository: libraryRepository,
	}
}
