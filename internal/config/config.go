package config

import (
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Build information -ldflags .
const (
	version    string = "dev"
	commitHash string = "-"
)

var cfg *Config

// GetConfigInstance returns service config
func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}

type RtgService struct {
	Address string `yaml:"address"`
}

type RtgFacade struct {
	Address string `yaml:"address"`
}

// Retry - contains all parameters of reconnect and resend.
type Retry struct {
	Count int           `yaml:"count"`
	Delay time.Duration `yaml:"delay"`
}

// Project - contains all parameters project information.
type Project struct {
	Debug       bool   `yaml:"debug"`
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	ServiceName string `yaml:"servicename"`
	Version     string
	CommitHash  string
}

// Telemetry - contains parameters for log server
type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project    Project       `yaml:"project"`
	RtgService RtgService    `yaml:"rtg_service"`
	RtgFacade  RtgFacade     `yaml:"rtg_facade"`
	Timeout    time.Duration `yaml:"timeout"`
	Retry      Retry         `yaml:"retry"`
	Telemetry  Telemetry     `yaml:"telemetry"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.Project.Version = version
	cfg.Project.CommitHash = commitHash

	return nil
}
