package services

import (
	"github.com/gurame/tiger/app/core/common"
	interfaces "github.com/gurame/tiger/app/core/interfaces/repositories"
	"github.com/gurame/tiger/app/core/models"
)

type ISellerService interface {
	Add(req models.SellerCreateReq) (models.SellerCreateRes, error)
}

type SellerService struct {
	repository interfaces.ISellerRepository
}

func NewSellerService(r interfaces.ISellerRepository) ISellerService {
	return &SellerService{
		repository: r,
	}
}

func (s *SellerService) Add(req models.SellerCreateReq) (models.SellerCreateRes, error) {

	errors := common.Validate(req)
	if errors != nil {
		panic(errors)
	}

	res, err := s.repository.Add(req)
	return res, err
}
