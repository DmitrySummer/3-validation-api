package readfiles

import (
	"fmt"
	"os"
	"strings"
)

// Функция для прочтения файла
func ReadFile(fileName string, hash string) bool {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return false
	}

	return strings.Contains(string(fileContent), hash)
}
