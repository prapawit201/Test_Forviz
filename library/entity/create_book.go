package entity

import "github.com/google/uuid"

type CreateBookRequest struct {
	Name   string  `json:"name" validate:"required,max=150"`
	Author string  `json:"author" validate:"required,max=100"`
	Type   string  `json:"type" validate:"required,max=50"`
	Price  float64 `json:"price" validate:"required"`
	Zone   *string `json:"zone"`
}

type CreateBookResponse struct{}

func (b *CreateBookRequest) ToEntityCreate() *Book {
	return &Book{
		UUID:   uuid.New().String(),
		Name:   b.Name,
		Author: b.Author,
		Type:   b.Type,
		Price:  b.Price,
		Zone:   b.Zone,
	}
}
