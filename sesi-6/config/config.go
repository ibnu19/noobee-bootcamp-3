package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App `yaml:"app"`
	DB  `yaml:"db"`
}

type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

var Cfg *Config

func LoadConfig(fileName string) (err error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	// cfg := Config{}
	var cfg Config
	log.Println(cfg)
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return
	}

	Cfg = &cfg
	return
}
