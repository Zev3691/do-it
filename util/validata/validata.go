package validata

import "github.com/go-playground/validator/v10"

var validateEnt *validator.Validate

func Init() {
	validateEnt = validator.New()
}

func Struct(s interface{}) error {
	return validateEnt.Struct(s)
}
