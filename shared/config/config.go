package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func NewConfig(target interface{}) error {
	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return err
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	if err := envconfig.Process("", target); err != nil {
		return err
	}

	return nil
}

type (
	Configuration struct {
		// - log config
		LogLevel     string `envconfig:"LOG_LEVEL" default:"INFO"`
		LogFilePath  string `envconfig:"LOG_FILE_PATH" default:"application.log"`
		LogFormatter string `envconfig:"LOG_FORMATTER" default:"TEXT"`
		LogMaxSize   int    `envconfig:"LOG_MAX_SIZE" default:"500"`
		LogMaxBackup int    `envconfig:"LOG_MAX_BACKUP" default:"3"`
		LogMaxAge    int    `envconfig:"LOG_MAX_BACKUP" default:"3"`
		LogCompress  bool   `envconfig:"LOG_COMPRESS" default:"true"`

		// - database
		DbUser                  string `envconfig:"DB_USER" required:"true" default:"root"`
		DbPass                  string `envconfig:"DB_PASS" required:"true" default:"root"`
		DbName                  string `envconfig:"DB_NAME" required:"true" default:"article"`
		DbHost                  string `envconfig:"DB_HOST" required:"true" default:"localhost"`
		DbPort                  int    `envconfig:"DB_PORT" required:"true" default:"5432"`
		DbMaxOpenConnection     int    `envconfig:"DB_MAX_OPEN_CONNECTION" required:"true" default:"32"`
		DbMaxIdleConnection     int    `envconfig:"DB_MAX_IDLE_CONNECTION" required:"true" default:"8"`
		DbMaxLifeTimeConnection string `envconfig:"DB_MAX_LIFE_TIME_CONNECTION" required:"true" default:"1h"`
		EnableLogSqlQuery       bool   `envconfig:"LOG_SQL_QUERY" default:"false"`
	}
)

func New() (*Configuration, error) {
	var (
		c = Configuration{}
	)

	if err := NewConfig(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
