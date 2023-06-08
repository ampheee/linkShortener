package postgresql

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"ozonFintech/config"
	"ozonFintech/pkg/logger"
	"ozonFintech/pkg/utilities"
	"time"
)

func GetPool(ctx context.Context, c config.Config) (connect *pgxpool.Pool, err error) {
	log := logger.GetLogger()
	err = utilities.ConnectWithTries(func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		connect, err = pgxpool.New(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			c.PostgreSQLDB.User,
			c.PostgreSQLDB.Pass,
			c.PostgreSQLDB.Host,
			c.PostgreSQLDB.Port,
			c.PostgreSQLDB.Dbname,
			c.PostgreSQLDB.SSLMode))
		return err
	}, 3, time.Second*3)
	if err != nil || connect.Ping(ctx) != nil {
		log.Fatal().Err(err).Msg("unable to connect to database. addr: " + fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			c.PostgreSQLDB.User,
			c.PostgreSQLDB.Pass,
			c.PostgreSQLDB.Host,
			c.PostgreSQLDB.Port,
			c.PostgreSQLDB.Dbname,
			c.PostgreSQLDB.SSLMode))
		return nil, err
	}
	log.Info().Msg("connected to database successfully")
	return connect, nil
}

func MigratesUp(c config.Config) {
	log := logger.GetLogger()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgreSQLDB.User,
		c.PostgreSQLDB.Pass,
		c.PostgreSQLDB.Host,
		c.PostgreSQLDB.Port,
		c.PostgreSQLDB.Dbname,
		c.PostgreSQLDB.SSLMode)
	m, err := migrate.New("file://../migrations/", dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get migrator.")
	}
	if err = m.Up(); err != nil {
		log.Warn().Err(err).Msg("Unable to up migrations.")
	}
	log.Info().Msg("Migration up done successfully!")
}

func MigratesDown(c config.Config) {
	log := logger.GetLogger()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgreSQLDB.User,
		c.PostgreSQLDB.Pass,
		c.PostgreSQLDB.Host,
		c.PostgreSQLDB.Port,
		c.PostgreSQLDB.Dbname,
		c.PostgreSQLDB.SSLMode)
	m, err := migrate.New("file://../migrations/", dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get migrator.")
	}
	if err = m.Down(); err != nil {
		log.Warn().Err(err).Msg("Unable to down migrations.")
	} else {
		log.Info().Msg("Migration down successfully!")
	}
}

func MigratesOnStart(c config.Config) {
	log := logger.GetLogger()
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgreSQLDB.User,
		c.PostgreSQLDB.Pass,
		c.PostgreSQLDB.Host,
		c.PostgreSQLDB.Port,
		c.PostgreSQLDB.Dbname,
		c.PostgreSQLDB.SSLMode)
	m, err := migrate.New("file://../migrations/", dbUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get migrator.")
	}
	if err = m.Migrate(1); err != nil && err != migrate.ErrNoChange {
		log.Warn().Err(err).Msg("Unable to rollback migrations.")
	} else {
		log.Info().Msg("Migration to start successfully!")
	}
}

func InitMigration(c config.Config) {
	switch c.Migration {
	case "Up":
		MigratesUp(c)
	case "Down":
		MigratesDown(c)
	case "OnStart":
		MigratesOnStart(c)
	default:
		return
	}
}
