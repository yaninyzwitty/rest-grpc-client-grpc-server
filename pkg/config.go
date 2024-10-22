package pkg

import (
	"io"
	"log/slog"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Srv `yaml:"server"`
	Database DB  `yaml:"database"`
}

type Srv struct {
	PORT int `yaml:"port"`
}

type DB struct {
	Path      string `yaml:"path"`
	Username  string `yaml:"username"`
	Passsword string `yaml:"password"`
}

func (c *Config) LoadConfig(file io.Reader) error {
	data, err := io.ReadAll(file)
	if err != nil {
		slog.Error("failed to read file", "error", err)
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		slog.Error("Failed to unmarshal file data", "error", err)
		return err
	}
	return nil

}
