package config

import (
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
	Docker   bool
}

type Server struct {
	Bind string `json:"bind"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

var conf = Config{
	Server{Bind: ":3001"},
	Database{User: "", Password: "", Address: ":5432"},
	false,
}

func Get() *Config {
	return &conf
}

func Parse(path string) error {
	buff, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(buff, &conf)
	if err != nil {
		return err
	}

	slog.With("conf", conf).Debug("")

	return nil
}
