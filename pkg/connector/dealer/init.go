package dealerconnector

import "models"

type impl struct {
	url string
}

func New(url models.DealerURL) DealerConnector {
	return &impl{url: url.String()}
}
