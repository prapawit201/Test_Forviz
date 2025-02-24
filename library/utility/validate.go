package utility

import (
	"log"

	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, input interface{}) error {
	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// validate error
			for _, err := range validationErrors {
				log.Printf("Field: '%s' failed validation rule: '%s'\n", err.Field(), err.ActualTag())
				return err
			}
		} else {
			log.Printf("Unexpected validation error: %s\n", err.Error())
			return err
		}
	}

	return nil
}
