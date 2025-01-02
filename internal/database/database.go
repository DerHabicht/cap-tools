package database

import (
	"database/sql"
	"fmt"

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
	// TODO: Use connection pools
	url := GetDBUrl()
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
