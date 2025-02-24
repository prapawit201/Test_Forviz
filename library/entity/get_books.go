package entity

import (
	"strings"

	"github.com/prapawit201/Test_Forviz/library/types/constant"
	"go.openly.dev/pointy"
)

type GetBooksRequest struct {
	Page        int     `json:"page" validate:"required,min=1"`
	Limit       int     `json:"limit" validate:"required,min=1"`
	FilterBy    *string `json:"filter_by" validate:"omitempty,oneof=name author type"`
	FilterValue *string `json:"filter_value"`
	Sort        *string `json:"sort" validate:"omitempty,oneof=id:asc id:desc borrow_count:asc borrow_count:desc"`
}

type GetBooksResponse struct {
	Page      int           `json:"page"`
	Limit     int           `json:"limit"`
	TotalPage int           `json:"total_page"`
	TotalData int           `json:"total_data"`
	Data      []GetBookData `json:"data"`
}

type GetBookData struct {
	ID          string  `json:"id"` // from uuid in db
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Type        string  `json:"type"`
	Zone        *string `json:"zone"`
	IsBorrow    bool    `json:"is_borrow"`
	BorrowCount int     `json:"borrow_count"`
}

func (g *GetBooksRequest) ToEntityWhereClauseBook() *Book {
	resp := Book{}

	if g.FilterBy != nil && g.FilterValue != nil {
		filterBy := pointy.StringValue(g.FilterBy, "")

		switch filterBy {
		case constant.WhereClauseName:
			resp.Name = pointy.StringValue(g.FilterValue, "")
		case constant.WhereClauseAuthor:
			resp.Author = pointy.StringValue(g.FilterValue, "")
		case constant.WhereClauseType:
			resp.Type = pointy.StringValue(g.FilterValue, "")
		default:
		}

	}
	return &resp
}

func (g *GetBooksRequest) ToEntityPaginate() BookPaginate {
	sort := "id asc" //default

	if g.Sort != nil {
		sort = strings.ReplaceAll(pointy.StringValue(g.Sort, ""), ":", " ")
	}

	return BookPaginate{
		Page:        g.Page,
		Limit:       g.Limit,
		FilterBy:    g.FilterBy,
		FilterValue: g.FilterValue,
		Sort:        sort,
	}
}

func (g *GetBooksResponse) ToEntityResponse(input *GetBooksRequest, bookCount int, data []Book) *GetBooksResponse {
	var (
		totalCount = 0
		totalPage  = 0
		resp       = make([]GetBookData, 0)
	)

	//not found
	if bookCount < 1 {
		totalCount = 0
		totalPage = 1
	} else {
		totalCount = bookCount
		totalPage = (bookCount + input.Limit - 1) / input.Limit

		for _, v := range data {
			result := GetBookData{
				ID:          v.UUID,
				Name:        v.Name,
				Author:      v.Author,
				Type:        v.Type,
				Zone:        v.Zone,
				IsBorrow:    v.IsBorrow,
				BorrowCount: v.BorrowCount,
			}
			resp = append(resp, result)
		}
	}

	return &GetBooksResponse{
		Page:      input.Page,
		Limit:     input.Limit,
		TotalPage: totalPage,
		TotalData: totalCount,
		Data:      resp,
	}
}
