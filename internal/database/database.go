package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
)

func GetDBUrl() string {
	_ = viper.ConfigFileUsed()

	if viper.GetBool("database.ssl") {
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			viper.GetString(config.DatabaseUser),
			viper.GetString(config.DatabasePassword),
			viper.GetString(config.DatabaseHost),
			viper.GetString(config.DatabasePort),
			viper.GetString(config.DatabaseName),
		)
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString(config.DatabaseUser),
		viper.GetString(config.DatabasePassword),
		viper.GetString(config.DatabaseHost),
		viper.GetString(config.DatabasePort),
		viper.GetString(config.DatabaseName),
	)
}

func GetDB() (*sql.DB, error) {
	url := GetDBUrl()
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

func DBVersion() (int64, error) {
	db, err := GetDB()
	if err != nil {
		return -1, errors.WithStack(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT version FROM schema_migrations;")
	if err != nil {
		return -1, errors.WithStack(err)
	}
	defer rows.Close()

	rows.Next()
	var version int64
	err = rows.Scan(&version)
	if err != nil {
		if strings.Contains(err.Error(), "Rows are closed") {
			return 0, nil
		}
		return -1, errors.WithStack(err)
	}

	return version, nil
}
