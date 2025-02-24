package entity

type DeleteBookRequest struct {
	ID string `json:"id" validate:"required"` // use for request/response is field uuid in DB
}
