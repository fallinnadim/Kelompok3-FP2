package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}

func ParseError(err error) (int, []string) {
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
		return http.StatusBadRequest, errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok { // UnmarshalTypeError saat shouldBindJSON, memastikan tipe data yang diisi harus sesuai dengan struct
		return http.StatusBadRequest, []string{fmt.Sprintf("Key %s harus bernilai %s", marshallingErr.Field, marshallingErr.Type.String())}
	} else if re, ok := err.(*RequestError); ok { // Custom error
		return re.StatusCode, []string{re.Err.Error()}
	}
	return http.StatusInternalServerError, []string{err.Error()} // kasus lain
}
