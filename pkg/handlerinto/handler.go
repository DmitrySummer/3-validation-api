package handlerinto

import (
	"3-validation-api/pkg/createjson"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Функция по обработке входящего запроса
func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		createjson.Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		createjson.Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}

// Функция по проверке правильного ввода пользователя
func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}

// Функция по декодированию поступающего запроса
func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
