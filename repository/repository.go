// This file contains the repository implementation layer.
package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	database, err := gorm.Open(postgres.Open(opts.Dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}

	sqlDB, err := database.DB()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(1 * time.Minute)

	dbMigration(opts)

	return &Repository{database}
}

func dbMigration(opts NewRepositoryOptions) {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}
	driver, err := migratePostgres.WithInstance(db, &migratePostgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./repository/migrations",
		"postgres", driver)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}
	m.Close()
	driver.Close()
	db.Close()
}
