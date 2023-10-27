package database

import (
	"fmt"
	"sesi-6/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM(config config.DB) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Pass, config.Name, config.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	conn, err := db.DB()
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxIdleTime(5 * time.Second)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxLifetime(time.Hour)
	return
}
