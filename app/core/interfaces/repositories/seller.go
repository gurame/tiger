package interfaces

import "github.com/gurame/tiger/app/core/models"

type ISellerRepository interface {
	Add(book models.SellerCreateReq) (models.SellerCreateRes, error)
}
