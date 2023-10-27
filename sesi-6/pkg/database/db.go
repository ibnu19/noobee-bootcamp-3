package database

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type ConnDB struct {
	Gorm *gorm.DB
	SqlX *sqlx.DB
}
