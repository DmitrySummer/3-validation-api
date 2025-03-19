package verifi

import (
	"3-validation-api/config"
	"3-validation-api/pkg/createjson"
	"3-validation-api/pkg/gethash"
	"3-validation-api/pkg/handlerinto"
	"3-validation-api/pkg/readfiles"
	"3-validation-api/pkg/sendemail"
	"fmt"
	"net/http"
)

type VerifiHandler struct {
	*config.Config
}

type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type VerifyHash struct {
	Email string `json:"email" validate:"required,email"`
	Hash  string `json: "hash" validate:"required,hash"`
}

func NewAuthHandler(router *http.ServeMux, deps VerifiHandler) {
	handler := &VerifiHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("POST /verify/{hash}", handler.Verify())
}

// Обработчик http запроса POST /send
func (handler *VerifiHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := handlerinto.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "123",
		}
		userHash := gethash.GetHash(body.Email)

		err = createjson.WriteJSON("UserEmail+Hash.txt", body.Email, userHash)
		if err != nil {
			fmt.Printf("Ошибка записи в файл ", err)
		}
		userConfig := config.LoadConfig()

		err = sendemail.SendEmail(userConfig, body.Email, userHash)
		if err != nil {
			fmt.Printf("Ошибка отправки email: ", err)
		}
		createjson.Json(w, data, 200)
	}
}

// Обработчик http запроса POST /verify/{hash}
func (handler *VerifiHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := handlerinto.HandleBody[VerifyHash](&w, r)
		if err != nil {
			return
		}

		if readfiles.ReadFile("UserEmail+Hash.txt", body.Hash) {
			createjson.Json(w, "Email подтвержден", http.StatusOK)
		} else {
			createjson.Json(w, "Неверный код подтверждения", http.StatusBadRequest)
		}
	}
}
