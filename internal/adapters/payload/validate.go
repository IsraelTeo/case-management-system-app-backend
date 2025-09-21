package payload

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func Init() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}
