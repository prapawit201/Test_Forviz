package entity

import "github.com/prapawit201/Test_Forviz/library/types/constant"

type UpdateBookStatusRequest struct {
	ID     string `json:"id" validate:"required"`                          // uuid
	Status string `json:"status" validate:"required,oneof=borrow receive"` // borrow ยืม, receive = ไม่ได้ยืม / user มาคืนเเล้ว

}

type UpdateBookStatusResponse struct{}

func (b *UpdateBookStatusRequest) ToEntityCreateBookBorrowLog(bookID int) *BookBorrowLog {

	return &BookBorrowLog{
		BookID: bookID,
		Status: b.Status,
	}
}

func (u *UpdateBookStatusRequest) ToEntityUpdate(bookData *Book) *Book {
	resp := Book{}

	//uuid
	resp.UUID = u.ID

	//status
	if u.Status == constant.BookStatusBorrow {
		resp.IsBorrow = true
		resp.BorrowCount = bookData.BorrowCount + 1

	} else if u.Status == constant.BookStatusReceive {
		resp.IsBorrow = false
	}

	return &resp
}
