package config

import (
	"fmt"
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

// Database - contains all parameters database connection.
type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
	Driver     string `yaml:"driver"`
}

// Grpc - contains parameter address grpc.
type Grpc struct {
	Port              int    `yaml:"port"`
	MaxConnectionIdle int64  `yaml:"maxConnectionIdle"`
	Timeout           int64  `yaml:"timeout"`
	MaxConnectionAge  int64  `yaml:"maxConnectionAge"`
	Host              string `yaml:"host"`
}

// Swagger - contains parameters for swagger port
type Swagger struct {
	Path     string `yaml:"path"`
	Filepath string `yaml:"filepath"`
}

// Router - contains parameters router port.
type Router struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Rest - contains parameter rest json connection.
type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
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

// Metrics - contains all parameters metrics information.
type Metrics struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

// Jaeger - contains all parameters metrics information.
type Jaeger struct {
	Service string `yaml:"service"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

// Kafka - contains all parameters kafka information.
type Kafka struct {
	Capacity uint64   `yaml:"capacity"`
	Topic    string   `yaml:"topic"`
	GroupID  string   `yaml:"groupId"`
	Brokers  []string `yaml:"brokers"`
}

// Status config for service.
type Status struct {
	Port          int    `yaml:"port"`
	Host          string `yaml:"host"`
	VersionPath   string `yaml:"versionPath"`
	LivenessPath  string `yaml:"livenessPath"`
	ReadinessPath string `yaml:"readinessPath"`
}

// Gateway - contains parameters for grpc-gateway port
type Gateway struct {
	Port               int      `yaml:"port"`
	Host               string   `yaml:"host"`
	Path               string   `yaml:"path"`
	AllowedCORSOrigins []string `yaml:"allowedCorsOrigins"`
}

// Telemetry - contains parameters for log server
type Telemetry struct {
	GraylogPath string `yaml:"graylogPath"`
}

// Retranslator is parameters of retranslator
type Retranslator struct {
	ChannelSize uint64 `yaml:"channelsize"`

	ConsumerCount  uint64        `yaml:"consumercount"`
	ConsumeSize    uint64        `yaml:"consumesize"`
	ConsumeTimeout time.Duration `yaml:"consumetimeout"`

	ProducerCount uint64 `yaml:"producercount"`
	WorkerCount   int    `yaml:"workercount"`

	WorkerPoolTimeout time.Duration `yaml:"workerpooltimeout"`
	WorkerQueueLen    int           `yaml:"workerqueuelen"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project      Project      `yaml:"project"`
	Grpc         Grpc         `yaml:"grpc"`
	Gateway      Gateway      `yaml:"gateway"`
	Swagger      Swagger      `yaml:"swagger"`
	Router       Router       `yaml:"router"`
	Rest         Rest         `yaml:"rest"`
	Database     Database     `yaml:"database"`
	Metrics      Metrics      `yaml:"metrics"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	Kafka        Kafka        `yaml:"kafka"`
	Status       Status       `yaml:"status"`
	Telemetry    Telemetry    `yaml:"telemetry"`
	Retranslator Retranslator `yaml:"retranslator"`
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

//DSN is return dsn string
func (db Database) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Name) //jdbc:postgresql://localhost:5432/rtg_service_api
}
