package repository

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/entity"
	"github.com/prapawit201/Test_Forviz/library/repository/model"
	"github.com/prapawit201/Test_Forviz/library/types/constant"
	"go.openly.dev/pointy"
)

// GetBooks implements libraryRepository.
func (u *libraryRepository) GetBooks(ctx *fiber.Ctx, input *entity.Book, paginate ...entity.BookPaginate) (int, []entity.Book, error) {
	var (
		db         = u.DB
		tx         = u.DB
		books      model.BookArray
		totalCount int64
	)

	// for where clause
	if input != nil {
		whereClause := new(model.Book).ToModel(input)
		db = db.Where(&whereClause)
	}

	db = db.Model(&model.Book{}).Count(&totalCount)
	if db.Error != nil {
		log.Printf("Repositorty CountBooks error: %s", db.Error.Error())
		return 0, nil, db.Error
	}

	if len(paginate) > 0 {
		offset := (paginate[0].Page - 1) * paginate[0].Limit

		if offset > int(totalCount) {
			offset = int(totalCount)
		}

		tx = u.DB.Offset(offset).Limit(paginate[0].Limit)

		//paginate
		if paginate[0].FilterBy != nil && paginate[0].FilterValue != nil {
			var (
				filterBy    = pointy.StringValue(paginate[0].FilterBy, "")
				filterValue = pointy.StringValue(paginate[0].FilterValue, "")
				whereClause = ""
			)

			switch filterBy {
			case constant.WhereClauseName:
				whereClause = string(constant.FilterByName)
			case constant.WhereClauseAuthor:
				whereClause = string(constant.FilterByAuthor)
			case constant.WhereClauseType:
				whereClause = string(constant.FilterByType)
			default:
			}

			// where condition for query data
			tx = tx.Where(whereClause, filterValue)
		}

		//sort
		tx = tx.Order(paginate[0].Sort)
	}

	tx = tx.Find(&books)
	if tx.Error != nil {
		log.Printf("Repositorty GetBooks error: %s", tx.Error.Error())
		return 0, nil, tx.Error
	}

	return int(totalCount), books.ToArrayEntity(), nil
}
