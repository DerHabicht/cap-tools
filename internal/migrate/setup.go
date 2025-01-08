package migrate

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
	"github.com/derhabicht/cap-tools/internal/database"
)

func rootDBURL() string {
	_ = viper.ConfigFileUsed()

	if viper.GetBool("database.ssl") {
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			rootDBUser,
			rootDBPassword,
			viper.GetString(config.DatabaseHost),
			viper.GetString(config.DatabasePort),
			"postgres",
		)
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		rootDBUser,
		rootDBPassword,
		viper.GetString(config.DatabaseHost),
		viper.GetString(config.DatabasePort),
		"postgres",
	)
}

func rootDB() (*sql.DB, error) {
	url := rootDBURL()
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

func setup() (*migrate.Migrate, *polluter.Polluter, error) {
	err := create()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	dbURL := database.GetDBUrl()
	log.Debug().Str("url", dbURL).Msg("database connection created")

	migrationsPath := viper.GetString(config.DatabaseMigrationSource)
	log.Debug().Str("path", migrationsPath).Msg("migrations path identified")

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	db, err := database.GetDB()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	p := polluter.New(polluter.PostgresEngine(db), polluter.YAMLParser)

	return m, p, nil
}
