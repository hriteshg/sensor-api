package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migrationsql "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func OpenConn(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.Host, config.Port, config.Username, config.Name, config.Password)
	log.Infof("Establishing connection to database: %s", dsn)
	gormDB, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Get the underlying *sql.DB instance
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// Open the database connection
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Infof("Successfully established connection to database: %s", config.Name)
	return gormDB, nil
}

func RunMigrations(d *gorm.DB, c DatabaseConfig, migrationPath string) error {
	log.Infof("Running schema migration %s", migrationPath)
	db, err := d.DB()
	if err != nil {
		return err
	}

	drvr, err := migrationsql.WithInstance(db, &migrationsql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		c.Name, drvr)
	if err != nil {
		log.Errorf("Running schema migration NewWithDatabaseInstance returned error %v", err)
		return err
	}

	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Info("No schema changes to apply")
		return nil
	}

	return err
}
