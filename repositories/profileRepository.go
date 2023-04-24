package repositories

import (
	"log"

	"github.com/monitoring-consumer/database"
	"github.com/monitoring-consumer/models"
)

func GetProfile(id string) (profile models.Profile, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		log.Fatalf("Error on open database connection: %v", err)
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT username, email FROM profiles WHERE id = $1`, id)

	err = row.Scan(&profile.Username, &profile.Email)

	return
}
