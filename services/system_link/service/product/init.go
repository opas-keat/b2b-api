package product

import (
	"systemlink/repo/product"
)

type ServiceImpl struct {
	productRepo product.Repo
}

func New(
	productRepo product.Repo,
) (Service, error) {
	return &ServiceImpl{
		productRepo,
	}, nil
}
