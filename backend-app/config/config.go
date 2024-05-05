package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type db struct {
	User         string `json:"user" yaml:"user"`
	Pass         string `json:"pass" yaml:"pass"`
	DatabaseName string `json:"databaseName" yaml:"databaseName"`
}

type Config struct {
	Db db `json:"db" yaml:"db"`
}

func Init() (Config, error) {
	var config Config
	f, err := os.Open("./config/config.yml")
	if err != nil {
		return config, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
