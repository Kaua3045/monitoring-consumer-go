package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/monitoring-consumer/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		log.Fatalf("Error on open database connection: %v", err)
	}

	err = conn.Ping()

	return conn, err
}
