package helper

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error) []string {
	// ValidationError type
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			switch e.Tag() {
			case "required":
				errorMessages[i] = fmt.Sprintf("Field %s harus diberi nilai", e.Field())
			case "min":
				errorMessages[i] = fmt.Sprintf("Field %s harus memiliki panjang minimal %s", e.Field(), e.Param())
			case "email":
				errorMessages[i] = fmt.Sprintf("Field %s harus memiliki format %s", e.Field(), e.Field())
			case "gte":
				errorMessages[i] = fmt.Sprintf("Field %s harus memiliki nilai lebih dari %s", e.Field(), e.Param())
			}
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok { // UnmarshalTypeError saat shouldBindJSON, memastikan tipe data yang diisi harus sesuai dengan struct
		return []string{fmt.Sprintf("Key %s harus bernilai %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return []string{err.Error()} // kasus lain
}
