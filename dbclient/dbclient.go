package dbclient

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Session  string `env:"DB_SESSION"`
	Table    string `env:"DB_NAME"`
}

func New(c *Config) *sql.DB {
	driverName := "mysql"
	dataSourceName := (*c).User + ":" + (*c).Password + "@tcp(" + (*c).Session + ":3306)/" + (*c).Table + "?charset=utf8"

	DB, connectionError := sql.Open(driverName, dataSourceName)
	if connectionError != nil {
		panic(connectionError.Error())
	}
	fmt.Println("You are now connected to a MySQL Database!")
	return DB
}
