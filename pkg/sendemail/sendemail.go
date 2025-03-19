package sendemail

import (
	"3-validation-api/config"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// Функция по отправке email сообщения пользователю
func SendEmail(config *config.Config, toEmail, hash string) error {
	verifyURL := fmt.Sprintf("http://localhost:8081/verify/%s", hash)
	message := fmt.Sprintf("Подтверждение Вашего Email\n\nПожалуйста, перейдите по ссылке для подтверждения: %s", verifyURL)

	e := email.NewEmail()
	e.From = "Письмо для подтверждения email <" + config.UserEmail + ">"
	e.To = []string{toEmail}
	e.Subject = "Подтверждение Email"
	e.Text = []byte(message)

	auth := smtp.PlainAuth("", config.UserEmail, config.UserPassword, config.UserHost)
	err := e.Send(config.UserHost+config.UserPort, auth)
	if err != nil {
		return fmt.Errorf("ошибка отправки email: %v", err)
	}

	fmt.Println("Письмо успешно отправлено!")
	return nil
}
