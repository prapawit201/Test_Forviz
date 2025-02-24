package http

import (
	"github.com/prapawit201/Test_Forviz/library/service"
)

type Handler struct {
	LibraryService service.LibraryService
}

func NewHTTPHandler(libraryService service.LibraryService) ServerInterface {
	return &Handler{
		LibraryService: libraryService,
	}
}
