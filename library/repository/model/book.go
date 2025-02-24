package model

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"gorm.io/gorm"
)

type Book struct {
	ID          int64   `json:"id" gorm:"primaryKey"`
	UUID        string  `json:"uuid" gorm:"uuid"`
	Name        string  `json:"name" gorm:"name"`
	Author      string  `json:"author" gorm:"author"`
	Type        string  `json:"type" gorm:"type"`
	Price       float64 `json:"price" gorm:"price"`
	Zone        *string `json:"zone" gorm:"zone"`
	BorrowCount int     `json:"borrow_count" gorm:"borrow_count"`
	IsBorrow    bool    `json:"is_borrow" gorm:"is_borrow"`
	DeletedAt   gorm.DeletedAt
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookArray []Book

func (b *Book) TableName() string {
	return "books"
}

func (b *Book) ToModel(input *entity.Book) *Book {
	resp := Book{}
	_ = copier.Copy(&resp, input)

	return &resp
}

func (b *Book) ToModelCreate(input *entity.Book) *Book {
	resp := Book{}
	_ = copier.Copy(&resp, input)

	return &resp
}

func (b *Book) ToModelUpdate(input *entity.Book) *Book {
	resp := Book{}
	_ = copier.Copy(&resp, input)
	return &resp
}

func (b *Book) ToEntity() *entity.Book {
	resp := entity.Book{}
	_ = copier.Copy(&resp, b)

	resp.CreatedAt = b.CreatedAt.String()
	resp.UpdatedAt = b.UpdatedAt.String()
	return &resp
}

func (b BookArray) ToArrayEntity() []entity.Book {
	resp := []entity.Book{}

	for _, v := range b {
		book := entity.Book{}
		_ = copier.Copy(&book, v)

		book.CreatedAt = v.CreatedAt.String()
		book.UpdatedAt = v.UpdatedAt.String()

		resp = append(resp, book)
	}
	return resp
}
