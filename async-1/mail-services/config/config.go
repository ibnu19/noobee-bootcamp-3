package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App   App         `yaml:"app"`
	Email EmailConfig `yaml:"email"`
}

type App struct {
	Port string `yaml:"port"`
}

type EmailConfig struct {
	SmtpHost     string `yaml:"config_smtp_host"`
	SmtpPort     int    `yaml:"config_smtp_port"`
	SenderName   string `yaml:"config_sender_name"`
	AuthEmail    string `yaml:"config_auth_email"`
	AuthPassword string `yaml:"config_auth_password"`
}

func LoadConfig() (config Config, err error) {
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &config)
	return
}
