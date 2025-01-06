package pgrepo

import (
	"database/sql"
	"fmt"
)

type PGRepository struct {
	db *sql.DB
}

func NewPGRepository(db *sql.DB) *PGRepository {
	return &PGRepository{
		db: db,
	}
}

func (repo *PGRepository) Close() error {
	return repo.db.Close()
}

type ErrNotFound struct {
	Object string
	Key    any
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("%s identified by %v not found", err.Object, err.Key)
}
