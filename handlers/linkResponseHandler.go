package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/monitoring-consumer/configs"
	"github.com/monitoring-consumer/models"
	"github.com/monitoring-consumer/repositories"
)

func SaveUrlResponse(urlId string, requestTime int, statusCode int, statusText string) {
	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		log.Fatalf("Error on get America/Sao_Paulo location: %v", err)
	}

	urlResponse := models.LinkResponses{
		Id:                 uuid.New().String(),
		ResponseMessage:    statusText,
		ResponseStatusCode: statusCode,
		RequestTime:        int64(requestTime),
		VerifiedDate:       time.Now().In(location),
		LinkId:             urlId,
	}
	repositories.InsertLinkResponse(urlResponse)
}

func VerifyUrl(url string) (int, int, string) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		return 0, 500, "Internal Server Error"
	}

	defer response.Body.Close()

	statusTextParts := strings.Split(response.Status, " ")
	var statusText string

	if len(statusTextParts) > 1 {
		statusText = statusTextParts[1]
	}

	return int(time.Since(start).Milliseconds()), response.StatusCode, statusText
}

func SendInternalErrorMail(profileId string, title string) {
	// faz parse do template (carregando as variaveis, para configurar depois)
	template, err := template.ParseFiles("./templates/internalErrorEmailTemplate.html")

	if err != nil {
		log.Printf("Error on send internal error mail: %v", err)
	}

	// cria uma variavel com o tipo Buffer
	var mailBody bytes.Buffer

	profile, err := repositories.GetProfile(profileId)

	if err != nil {
		log.Printf("Error on get profile: %v", err)
	}

	location, err := time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		log.Fatalf("Error on get America/Sao_Paulo location: %v", err)
	}

	data := struct {
		ErrorTime string
		Title     string
	}{
		ErrorTime: time.Now().In(location).Format("02/01/2006 ás 15:04:05"),
		Title:     title,
	}

	// Passa o MailBody que vai receber a versão com os dados do usuário, que são passados no data
	err = template.Execute(&mailBody, data)

	if err != nil {
		log.Printf("Error on parse values to mail template: %v", err)
	}

	cfg := configs.GetMailConfig()

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.SmtpHost)

	mailTo := []string{profile.Email}

	message := []byte("To: " + mailTo[0] + "\r\n" +
		"Subject: " + title + " - Erro" + "\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n" +
		mailBody.String())

	err = smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.From, mailTo, message)

	if err != nil {
		log.Printf("Error on send mail: %v", err)
	}
}
