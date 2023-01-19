package shipping

import "systemlink/repo/shipping"

type ServiceImpl struct {
	shippingRepo shipping.Repo
}

func New(
	shippingRepo shipping.Repo,
) (Service, error) {
	return &ServiceImpl{
		shippingRepo,
	}, nil
}
