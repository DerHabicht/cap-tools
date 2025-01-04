package gormrepo

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type model struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *sql.DB) (*GormRepository, error) {
	gormCfg := &gorm.Config{}
	gormPGCfg := postgres.Config{Conn: db}
	dialector := postgres.New(gormPGCfg)

	gdb, err := gorm.Open(dialector, gormCfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &GormRepository{db: gdb}, nil
}
