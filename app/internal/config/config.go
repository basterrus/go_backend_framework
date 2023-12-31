package config

import (
	"flag"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	IsDebug       bool `yaml:"is-debug" env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `yaml:"is-development" env:"IS_DEV" env-default:"false"`
	MaxAttempts   int  `yaml:"max_attempts" env:"MAX_ATTEMPTS" env-default:"5"`
	MaxDelay      int  `yaml:"max_delay" env:"MAX_DELAY" env-default:"3"`
	IsBinary      bool `yaml:"is_binary" env:"IS_BINARY" env-default:"false"`
	Listen        struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	}
	HTTP struct {
		IP           string        `yaml:"ip" env:"HTTP-IP"`
		Port         int           `yaml:"port" env:"HTTP-PORT"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP-READ-TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP-WRITE-TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed_origins"`
			AllowCredentials   bool     `yaml:"allow_credentials"`
			AllowedHeaders     []string `yaml:"allowed_headers"`
			OptionsPassthrough bool     `yaml:"options_passthrough"`
			ExposedHeaders     []string `yaml:"exposed_headers"`
			Debug              bool     `yaml:"debug"`
		} `yaml:"cors"`
	} `yaml:"http"`
	GRPC struct {
		IP   string `yaml:"ip" env:"GRPC-IP"`
		Port int    `yaml:"port" env:"GRPC-PORT"`
	} `yaml:"grpc"`
	AppConfig struct {
		LogLevel  string `yaml:"log-level" env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `yaml:"email" env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `yaml:"password" env:"ADMIN_PWD" env-default:"admin"`
		} `yaml:"admin"`
	} `yaml:"app"`
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
	//MongoDB struct {
	//	Host               string        `yaml:"host" env-required:"true"`
	//	Port               string        `yaml:"port" env-required:"true"`
	//	Username           string        `yaml:"username"`
	//	Password           string        `yaml:"password"`
	//	AuthDB             string        `yaml:"auth_db" env-required:"true"`
	//	Database           string        `yaml:"database" env-required:"true"`
	//	Collection         string        `yaml:"collection" env-required:"true"`
	//	RequestContextTime time.Duration `yaml:"request_context_time" env-required:"true"`
	//} `yaml:"mongodb" env-required:"true"`
}

const (
	EnvConfigPath      = "CONFIG-PATH"
	FlagConfigPath     = "config"
	FlagConfigPathName = "/Users/pavelnedosivin/GolandProjects/go_backend_framework/config.yml"
	DescriptionText    = "this is app config file"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig(logger logging.Logger) *Config {
	once.Do(func() {
		flag.StringVar(
			&configPath,
			FlagConfigPath,
			FlagConfigPathName,
			DescriptionText,
		)
		flag.Parse()

		logger.Debug("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPath)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "The backend framework help text"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
