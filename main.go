package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/joho/godotenv"
	"github.com/monitoring-consumer/configs"
	"github.com/monitoring-consumer/handlers"
	"github.com/monitoring-consumer/models"
)

// {"id": "d18f7b89-601b-4d75-92d1-44b76bb1bed8", "owner_id": "b8ece0fa-5402-459f-bb39-9222fc1775aa", "title": "API - Monitoring", "url": "https://httpstat.us/500"}

func main() {
	// godotenv.Load()
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error on load configs: %v", err)
	}

	processMessages()
}

func processMessages() {
	cfg := configs.GetAWSConfig()

	// configure aws session
	config, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("Error on create aws session: %v", err)
	}

	// Cria o cliente do SQS
	sqsClient := sqs.NewFromConfig(config)

	// Url da fila

	// Receber as mensagens da fila
	for {
		result, err := sqsClient.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(cfg.SQSQueueUrl),
			MaxNumberOfMessages: *aws.Int32(10),
			VisibilityTimeout:   *aws.Int32(10),
			WaitTimeSeconds:     *aws.Int32(5),
		})

		if err != nil {
			log.Printf("Error on receive messages: %v", err)
			continue
		}

		// Processar as mensagens da fila
		for _, message := range result.Messages {

			// Converte o JSON para o Struct
			var linkStruct models.Link
			err := json.Unmarshal([]byte(*message.Body), &linkStruct)

			if err != nil {
				log.Printf("Error on decode mensagem: %v", err)
				continue
			}

			responseTime, statusCode, statusText := handlers.VerifyUrl(linkStruct.Url)

			// Salva no banco de dados a URL
			handlers.SaveUrlResponse(linkStruct.Id, responseTime, statusCode, statusText)

			// statusCodeStr := strconv.Itoa(statusCode)

			// if strings.HasPrefix(statusCodeStr, "5") {
			// 	go handlers.SendInternalErrorMail(linkStruct.Owner_id, linkStruct.Title)
			// }

			// Deleta a mensagem da fila
			_, err = sqsClient.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(cfg.SQSQueueUrl),
				ReceiptHandle: message.ReceiptHandle,
			})

			if err != nil {
				log.Printf("Error on delete message: %v", err)
			}
		}
	}
}
