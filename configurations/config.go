package configurations

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Server configuration
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	}
	// Database configuration
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	}
}

func (cfg *Config) Init() error {
	f, err := os.Open("config-dev.yaml")

	if err != nil {
		return err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)

	if err != nil {
		return err
	}
	return nil
}
