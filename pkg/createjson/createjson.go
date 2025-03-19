package createjson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Функция для записи файла JSON
func WriteJSON(fileName, email, hash string) error {
	data := fmt.Sprintf("%s,%s\n", email, hash)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

// Функция по добавлению в json
func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
