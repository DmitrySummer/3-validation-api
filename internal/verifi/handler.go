package verifi

import (
	"3-validation-api/config"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
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

func (handler *VerifiHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "123",
		}
		userHash := GetHash(body.Email)
		err = WriteJSON("UserEmail+Hash.txt", body.Email, userHash)
		if err != nil {
			fmt.Printf("Ошибка записи в файл ", err)
		}
		err = sendEmail(body.Email, userHash)
		if err != nil {
			fmt.Printf("Ошибка отправки email: ", err)
		}
		Json(w, data, 200)

	}

}

func (handler *VerifiHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := HandleBody[VerifyHash](&w, r)
		if err != nil {
			return
		}

		if ReadFile("UserEmail+Hash.txt", body.Hash) {
			Json(w, "Email подтвержден", http.StatusOK)
		} else {
			Json(w, "Неверный код подтверждения", http.StatusBadRequest)
		}
	}
}

func ReadFile(fileName string, hash string) bool {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return false
	}

	return strings.Contains(string(fileContent), hash)
}

func GetHash(toEmail string) string {
	hash := sha256.Sum256([]byte(toEmail))
	hexHash := hex.EncodeToString(hash[:])
	return hexHash
}

func sendEmail(toEmail, hash string) error {
	verifyURL := fmt.Sprintf("http://localhost:8081/verify/%s", hash)
	message := fmt.Sprintf("Подтверждение Вашего Email\n\nПожалуйста, перейдите по ссылке для подтверждения: %s", verifyURL)

	e := email.NewEmail()
	e.From = "Сервис подтверждения email <test@gmail.com>"
	e.To = []string{toEmail}
	e.Subject = "Подтверждение Email"
	e.Text = []byte(message)

	auth := smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com")
	err := e.Send("smtp.gmail.com:587", auth)
	if err != nil {
		return fmt.Errorf("ошибка отправки email: %v", err)
	}

	fmt.Println("Письмо успешно отправлено!")
	return nil
}

// Функция для записи файла JSON
func WriteJSON(fileName string, email string, hash string) error {
	data := fmt.Sprintf("%s,%s\n", email, hash)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		Json(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		Json(*w, err.Error(), 402)
		return nil, err
	}
	return &body, nil
}

func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
