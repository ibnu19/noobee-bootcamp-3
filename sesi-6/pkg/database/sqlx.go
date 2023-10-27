package database

import (
	"fmt"
	"sesi-6/config"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectSQLX(config config.DB) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Pass, config.Name,
	)

	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return
}
