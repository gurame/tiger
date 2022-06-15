package audit

import (
	"context"
	"time"

	dbcontext "github.com/gurame/tiger/app/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ConfigSeller() {
	dbcontext.AddSellerHook(boil.BeforeInsertHook, func(_ context.Context, _ boil.ContextExecutor, m *dbcontext.Seller) error {
		m.Createdby = "gurame"
		m.Created = time.Now()
		return nil
	})
}
