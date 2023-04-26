package repositories

import (
	"log"
	"time"

	"github.com/monitoring-consumer/database"
	"github.com/monitoring-consumer/models"
)

func InsertLinkResponse(response models.LinkResponses) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error on open database connection: %v", err)
	}
	defer conn.Close()

	sql := `INSERT INTO links_responses (id, response_message, status_code, verified_date, request_time, url_id)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = conn.Exec(sql,
		response.Id,
		response.ResponseMessage,
		response.ResponseStatusCode,
		response.VerifiedDate,
		response.RequestTime,
		response.LinkId)

	if err != nil {
		log.Fatalf("Error on insert link response in database: %v", err)
	}
}

func UpdatLink(linkId string, nextExecuteDate *time.Time) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error on open database connection: %v", err)
	}
	defer conn.Close()

	sql := `UPDATE links SET next_execute_date=$1 WHERE id=$2`

	_, err = conn.Exec(sql, nextExecuteDate, linkId)

	if err != nil {
		log.Fatalf("Error on update link next execute date in database: %v", err)
	}
}
