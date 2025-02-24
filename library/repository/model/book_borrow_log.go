package model

import (
	"time"

	"github.com/prapawit201/Test_Forviz/library/entity"
)

type BookBorrowLog struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	BookID    int    `json:"book_id" gorm:"book_id"`
	Status    string `json:"status" gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BookBorrowLog) TableName() string {
	return "book_borrow_logs"
}

func (b *BookBorrowLog) ToModel() *BookBorrowLog {
	return &BookBorrowLog{}
}

func (b *BookBorrowLog) ToModelCreate(input *entity.BookBorrowLog) *BookBorrowLog {
	return &BookBorrowLog{
		BookID: input.BookID,
		Status: input.Status,
	}
}
