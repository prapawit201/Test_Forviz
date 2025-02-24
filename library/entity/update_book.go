package entity

import "go.openly.dev/pointy"

type UpdateBookRequest struct {
	ID     string   `json:"id" validate:"required"` // uuid
	Name   *string  `json:"name" validate:"omitempty,max=150"`
	Author *string  `json:"author" validate:"omitempty,max=100"`
	Type   *string  `json:"type" validate:"omitempty,max=50"`
	Price  *float64 `json:"price"`
	Zone   *string  `json:"zone"`
}

type UpdateBookResponse struct{}

func (u *UpdateBookRequest) ToEntityUpdate() *Book {
	resp := Book{}

	//uuid
	resp.UUID = u.ID

	//name
	if u.Name != nil {
		resp.Name = pointy.StringValue(u.Name, "")
	}

	//author
	if u.Author != nil {
		resp.Author = pointy.StringValue(u.Author, "")
	}

	//type
	if u.Type != nil {
		resp.Type = pointy.StringValue(u.Type, "")
	}

	//price
	if u.Price != nil {
		resp.Price = pointy.Float64Value(u.Price, 0)
	}

	//zone
	if u.Zone != nil {
		resp.Zone = u.Zone
	}

	return &resp
}
