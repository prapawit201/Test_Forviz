package entity

type GetBookDetailRequest struct {
	ID string `json:"id" validate:"required"` // uuid
}

type GetBookDetailResponse struct {
	ID       string  `json:"id"` // from uuid in db
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Type     string  `json:"type"`
	Zone     *string `json:"zone"`
	IsBorrow bool    `json:"is_borrow"`
}

func (g *GetBookDetailResponse) ToEntity(input *Book) *GetBookDetailResponse {
	return &GetBookDetailResponse{
		ID:       input.UUID,
		Name:     input.Name,
		Author:   input.Author,
		Type:     input.Type,
		Zone:     input.Zone,
		IsBorrow: input.IsBorrow,
	}
}
