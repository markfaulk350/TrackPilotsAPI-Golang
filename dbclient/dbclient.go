package dbclient

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type Config struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Session  string `env:"DB_SESSION"`
	Table    string `env:"DB_NAME"`
}

func New(c *Config) *sql.DB {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	dataSourceName := (*c).User + ":" + (*c).Password + "@tcp(" + (*c).Session + ":3306)/" + (*c).Table + "?charset=utf8"

	DB, connectionError := sql.Open("mysql", dataSourceName)
	if connectionError != nil {
		logger.Error().Err(connectionError).Msg("Failed to connect to MySQL Database")
	}

	logger.Info().Msg("Successfuly connected to MySQL Database")
	return DB
}
