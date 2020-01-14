package main

import (
	"fmt"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"stu/sql_migrate/bdata"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
//go:generate go-bindata -prefix migrations -pkg bdata -o bdata/migrations_gen.go migrations
var g_db *sqlx.DB
var dns = "root:pPEgcZWi8q@tcp(localhost:3306)/test?parseTime=true&charset=utf8mb4&loc=Hongkong"
// OpenDatabase opens the database and performs a ping to make sure the
// database is up.
func OpenDatabase(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %s", err)
	}
	for {
		if err := db.Ping(); err != nil {
			log.Errorf("ping database error, will retry in 2s: %s", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}
	return db, nil
}

func setPostgreSQLConnection() error {
	log.Info("connecting to postgresql")
	db, err := OpenDatabase(dns)
	if err != nil {
		return errors.Wrap(err, "database connection error")
	}
	g_db = db
	return nil
}

func runDatabaseMigrations() error {
	log.Info("applying database migrations")
	m := &migrate.AssetMigrationSource{
		Asset:    bdata.Asset,
		AssetDir: bdata.AssetDir,
		Dir:      "",
	}
	n, err := migrate.Exec(g_db.DB, "mysql", m, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "applying migrations failed")
	}
	log.WithField("count", n).Info("migrations applied")

	return nil
}

func main() {
	setPostgreSQLConnection()
	err := runDatabaseMigrations()
	if err != nil {
		log.Error(err)
	}
}
