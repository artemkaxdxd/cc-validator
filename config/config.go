package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server Server
}

type Server struct {
	Port string `envconfig:"SERVER_PORT"`
}

var Cfg Config

func FillConfig() error {
	return envconfig.Process("", &Cfg)
}
