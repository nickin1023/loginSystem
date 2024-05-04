package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

type db struct {
	User         string `json:"user" yaml:"user"`
	Pass         string `json:"pass" yaml:"pass"`
	DatabaseName string `json:"databaseName" yaml:"databaseName"`
}

type config struct {
	Db db `json:"db" yaml:"db"`
}

var Config config

func init() {
	f, err := os.Open("./config/config.yml")
	if err != nil {
		slog.Error("config os.Open err:", err)
		os.Exit(1)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&Config)
	if err != nil {
		slog.Error("Yaml decode error:", err)
		os.Exit(1)
	}
}
