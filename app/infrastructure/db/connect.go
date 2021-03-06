package db

import (
	"database/sql"
	"log"

	"github.com/gurame/tiger/app/infrastructure/db/audit"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Connection() *sql.DB {
	db, err := sql.Open("postgres", "dbname=TigerDb user=admin password=admin sslmode=disable")
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	boil.SetDB(db)
	audit.ConfigSeller()

	return db
}
