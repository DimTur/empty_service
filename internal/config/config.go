package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer HTTPServer `yaml:"http_server"`
	Storage    Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:":8000"`
	Timeout      time.Duration `yaml:"timeout" env-default:"5s"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}

type Storage struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}
