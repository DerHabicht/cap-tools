package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/romanyx/polluter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
	_ "github.com/derhabicht/cap-tools/internal/config"
	"github.com/derhabicht/cap-tools/internal/database"
)

var (
	rootDBUser = "postgres"
	rootDBPass = "postgres"
)

func rootDBURL() string {
	_ = viper.ConfigFileUsed()

	if viper.GetBool("database.ssl") {
		return fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			rootDBUser,
			rootDBPass,
			viper.GetString(config.DatabaseHost),
			viper.GetString(config.DatabasePort),
			"postgres",
		)
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		rootDBUser,
		rootDBPass,
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

func create() error {
	_, err := database.DBVersion()
	if err != nil {
		e := errors.Cause(err)
		if !strings.Contains(
			e.Error(),
			fmt.Sprintf(`database "%s" does not exist`, viper.GetString(config.DatabaseName)),
		) {
			return errors.WithStack(err)
		}
	} else {
		return nil
	}

	db, err := rootDB()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("creating database")
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", viper.GetString(config.DatabaseName)))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func drop() error {
	db, err := rootDB()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("dropping database")
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", viper.GetString(config.DatabaseName)))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func reset(pollute bool) error {
	err := down()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		return errors.WithStack(err)
	}
	err = up(pollute)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func down() error {
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("migrating database down")
	m, _, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Down()
	if err != nil {
		if !strings.Contains(err.Error(), "no migration") {
			return errors.WithStack(err)
		}
	}

	version, dirty, err := m.Version()
	if err != nil {
		if strings.Contains(err.Error(), "no migration") {
			log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", 0).Bool("dirty",
				false).Msg("down migration complete")
			return nil
		}
		return errors.WithStack(err)
	}

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", version).Bool("dirty",
		dirty).Msg("down migration complete")
	return nil
}

func up(pollute bool) error {
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("migrating database up")
	m, p, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			return errors.WithStack(err)
		}
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", version).Bool("dirty",
		dirty).Msg("up migration complete")

	if pollute {
		err = seed(p)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func seed(polluter *polluter.Polluter) error {
	if polluter == nil {
		_, p, err := setup()
		if err != nil {
			return errors.WithStack(err)
		}
		polluter = p
	}

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("seeding database")
	seedPath := viper.GetString(config.DatabaseMigrationSeed)

	f, err := os.Open(seedPath)
	if err != nil {
		log.Error().Err(err).Str("database", viper.GetString("database.name")).Msg("unable to read database seed file")
	}
	defer f.Close()

	err = polluter.Pollute(f)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
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

func main() {
	pollute := flag.Bool("seed", false, "seed database with test data after migration")
	flag.StringVar(&rootDBUser, "psql-root-user", "postgres", "postgres admin username FOR DEV USE ONLY! ")
	flag.StringVar(&rootDBPass, "psql-root-password", "postgres", "postgres admin password FOR DEV USE ONLY! ")
	flag.Parse()

	cmd := flag.Arg(0)
	if cmd == "" {
		fmt.Println(
			"Usage: go run migrate.go " +
				"[--seed] [--psql-root-user=<postgres>] [--psql-root-password=<postgres>] " +
				"[create|drop|up|down|reset|seed]" +
				"<new migration name>",
		)
		os.Exit(1)
	}

	var err error
	switch strings.ToLower(cmd) {
	case "create":
		err = create()
	case "drop":
		err = drop()
	case "up":
		err = up(*pollute)
	case "down":
		err = down()
	case "reset":
		err = reset(*pollute)
	case "seed":
		err = seed(nil)
	default:
		fmt.Printf("Invalid command: %s\nValid commands are [up|down|reset]\n", cmd)
	}

	if err != nil {
		log.Error().Err(err).Str("command", cmd).Msg("problem encountered while running command")
		os.Exit(1)
	}

	os.Exit(0)
}
