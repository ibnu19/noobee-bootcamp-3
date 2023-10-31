package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewValidator() (validate *validator.Validate) {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fl reflect.StructField) string {
		name := strings.SplitN(fl.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return
}

func ValidatorForError(errs error) (errFields map[string]string) {
	errFields = make(map[string]string)
	for _, e := range errs.(validator.ValidationErrors) {
		errField := ""

		switch e.Tag() {
		case "required":
			errField = fmt.Sprintf("kolom %s tidak boleh kosong", e.Field())
		case "email":
			errField = "format email tidak valid"
		default:
			errField = fmt.Sprintf("%v", e.Error())
		}

		errFields[e.Field()] = errField
	}
	return
}
