package dbcon

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectToDB(cfg Config) (*pgxpool.Pool, error) {
	// Baglanti dizesi olusturuluyor
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	// 5 saniyelik zaman asimi ile bir context olusturuyoruz
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // fonskiyon bitince context'i iptal ediyoruz

	// Veritabani baglantisi icin pgxpool.Connect baglantisi cagriliyor
	db, err := pgxpool.Connect(ctx, dataSourceName)
	if err != nil {
		// Eger baglanti hatasi varsa, hatayi donduruyoruz
		return nil, err
	}

	// Her yeni baglamada veritabaninin calisip calismadigni kontrol ediyoruz
	if err = db.Ping(context.Background()); err != nil {
		return nil, errors.New("db.Ping")
	}

	return db, nil
}
