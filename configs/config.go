package configs

import "os"

var cfg *config

type config struct {
	DB        DBConfig
	AWSConfig AWSConfig
	Mail      MailConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type AWSConfig struct {
	SQSQueueUrl string
}

type MailConfig struct {
	From     string
	Username string
	Password string
	SmtpHost string
	SmtpPort string
}

func Load() (*config, error) {
	cfg = new(config)

	cfg.DB = DBConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
	}

	cfg.AWSConfig = AWSConfig{
		SQSQueueUrl: os.Getenv("QUEUE_URL"),
	}

	cfg.Mail = MailConfig{
		From:     os.Getenv("MAIL_FROM"),
		Username: os.Getenv("MAIL_USER"),
		Password: os.Getenv("MAIL_PASSWORD"),
		SmtpHost: os.Getenv("SMTP_HOST"),
		SmtpPort: os.Getenv("SMTP_PORT"),
	}

	return cfg, nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetAWSConfig() AWSConfig {
	return cfg.AWSConfig
}

func GetMailConfig() MailConfig {
	return cfg.Mail
}
