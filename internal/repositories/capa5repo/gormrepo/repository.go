package gormrepo

import (
	"database/sql"
)

type GormRepository struct {
	db *sql.DB
}
