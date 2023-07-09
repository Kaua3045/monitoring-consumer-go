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

type LinkExecutionType string
type LinkExecutionFunc func(*time.Time) *time.Time

const (
	NoRepeat       LinkExecutionType = "NO_REPEAT"
	EveryDay       LinkExecutionType = "EVERY_DAYS"
	OnSpecificDay  LinkExecutionType = "ON_SPECIFIC_DAY"
	TwoTimesAMonth LinkExecutionType = "TWO_TIMES_A_MONTH"
	EveryFiveHours LinkExecutionType = "EVERY_FIVE_HOURS"
	// Adicionar mais tipos de linkExecution aqui, se necessário
)

var linkExecutionFuncMap = map[LinkExecutionType]LinkExecutionFunc{
	NoRepeat: func(executeDate *time.Time) *time.Time {
		return nil
	},
	EveryDay: func(executeDate *time.Time) *time.Time {
		currentDay := time.Now().Day()
		nextExecuteDate := time.Date(
			executeDate.Year(),
			executeDate.Month(),
			currentDay+1,
			executeDate.Hour(),
			executeDate.Minute(),
			executeDate.Second(),
			executeDate.Nanosecond(),
			executeDate.Location())
		return &nextExecuteDate
	},
	OnSpecificDay: func(executeDate *time.Time) *time.Time {
		currentMonth := time.Now().Month()
		nextExecuteDate := time.Date(
			executeDate.Year(),
			currentMonth+1,
			executeDate.Day(),
			executeDate.Hour(),
			executeDate.Minute(),
			executeDate.Second(),
			executeDate.Nanosecond(),
			executeDate.Location())
		return &nextExecuteDate
	},
	TwoTimesAMonth: func(executeDate *time.Time) *time.Time {
		currentDay := time.Now().Day()
		nextExecuteDate := time.Date(
			executeDate.Year(),
			executeDate.Month(),
			currentDay+15,
			executeDate.Hour(),
			executeDate.Minute(),
			executeDate.Second(),
			executeDate.Nanosecond(),
			executeDate.Location())

		// TODO: aqui verifica se o mês atual é maior que o mês que veio, se for atualiza
		if time.Now().Local().Month() > executeDate.Month() {
			nextExecuteDate = nextExecuteDate.AddDate(0, 1, 0)
		}

		return &nextExecuteDate
	},
}

func UpdateUrl(urlId string, linkExecution string, executeDate string) {
	dateLayout := "2006-01-02T15:04:05Z"
	executeDateConverted, err := time.Parse(dateLayout, executeDate)

	if err != nil {
		log.Fatal("Error on parse date to specified layout")
	}

	getNextExecuteDate, ok := linkExecutionFuncMap[LinkExecutionType(linkExecution)]

	if !ok {
		log.Fatal("Error getting link execution type: this link execution type does not exist.")
	}

	repositories.UpdatLink(urlId, getNextExecuteDate(&executeDateConverted))
}
