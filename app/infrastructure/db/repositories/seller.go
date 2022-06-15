package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	interfaces "github.com/gurame/tiger/app/core/interfaces/repositories"
	"github.com/gurame/tiger/app/core/models"
	dbcontext "github.com/gurame/tiger/app/infrastructure/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SellerRepository struct {
	Db *sql.DB
}

//NewRepo is the single instance repo that is being created.
func NewSellerRepository(db *sql.DB) interfaces.ISellerRepository {
	return &SellerRepository{
		Db: db,
	}
}

//CreateBook is a mongo repository that helps to create books
func (r *SellerRepository) Add(req models.SellerCreateReq) (models.SellerCreateRes, error) {

	id := uuid.NewString()
	s := &dbcontext.Seller{Sellerid: id, Name: req.Name, Taxid: req.TaxId}
	s.Insert(context.Background(), r.Db, boil.Infer())

	res := models.SellerCreateRes{SellerId: id, Name: req.Name, TaxId: req.TaxId}

	return res, nil
}
