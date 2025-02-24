package utility

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prapawit201/Test_Forviz/library/types/constant"
)

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponseWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(c *fiber.Ctx, code int, err error, data ...interface{}) error {

	if err != nil {
		if err.Error() == constant.ErrorRecordNotFound {
			code = http.StatusNotFound
		}

		c.JSON(ErrorResponse{Code: code, Message: err.Error()})
	}

	if code == http.StatusOK {
		if len(data) > 0 {
			c.JSON(SuccessResponseWithData{Code: code, Message: constant.ResponseSuccess, Data: data[0]})
		} else {
			c.JSON(SuccessResponse{Code: code, Message: constant.ResponseSuccess})
		}

	} else if code == http.StatusCreated {
		c.JSON(SuccessResponse{Code: code, Message: constant.ResponseCreatedSuccess})
	}

	c.Status(code)
	return nil
}
