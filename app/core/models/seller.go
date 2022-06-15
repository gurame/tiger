package models

type (
	SellerCreateReq struct {
		TaxId string `json:"taxId" validate:"required"`
		Name  string `json:"name" validate:"required"`
	}

	SellerCreateRes struct {
		SellerId string
		TaxId    string
		Name     string
	}
)
