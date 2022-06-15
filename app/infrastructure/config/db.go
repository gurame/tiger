package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gurame/tiger/app/infrastructure/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Database() {
	database, err := sql.Open("postgres", "dbname=TigerDb user=admin password=admin sslmode=disable")
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database connection success!")
	boil.SetDB(database)

	db.AddSellerHook(boil.BeforeInsertHook, func(_ context.Context, _ boil.ContextExecutor, m *db.Seller) error {
		m.Createdby = "gurame"
		m.Created = time.Now()
		return nil
	})
}
