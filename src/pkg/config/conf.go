package config

import (
	"gopkg.in/yaml.v3"
	"os"

	"github.com/caarlos0/env/v10"
)

type ServerConf struct {
	HttpConf `yaml:"http"`
	MysqlConf
}

type HttpConf struct {
	Port int `yaml:"port"`
}

type MysqlConf struct {
	User         string `env:"MYSQL_USER"`
	Passward     string `env:"MYSQL_PASSWORD"`
	Host         string `env:"MYSQL_HOST"`
	Port         string `env:"MYSQL_PORT"`
	DatabaseName string `env:"MYSQL_DATABASE"`
}

func LoadConf(path string) (*ServerConf, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := ServerConf{}
	yaml.Unmarshal(b, &conf)

	if err := env.Parse(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
