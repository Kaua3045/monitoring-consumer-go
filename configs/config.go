package configs

import "github.com/spf13/viper"

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

func Load() error {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.DB = DBConfig{
		Host:     viper.GetString("PG_HOST"),
		Port:     viper.GetString("PG_PORT"),
		User:     viper.GetString("PG_USER"),
		Password: viper.GetString("PG_PASSWORD"),
		Database: viper.GetString("PG_DATABASE"),
	}

	cfg.AWSConfig = AWSConfig{
		SQSQueueUrl: viper.GetString("QUEUE_URL"),
	}

	cfg.Mail = MailConfig{
		From:     viper.GetString("MAIL_FROM"),
		Username: viper.GetString("MAIL_USER"),
		Password: viper.GetString("MAIL_PASSWORD"),
		SmtpHost: viper.GetString("SMTP_HOST"),
		SmtpPort: viper.GetString("SMTP_PORT"),
	}

	return nil
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
