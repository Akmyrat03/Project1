package dbcon

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectToDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host = %s port = %s user = %s password = %s dbname = %s sslmode = %s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	DB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, err
}
