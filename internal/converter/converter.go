package converter

import (
	"net/http"

	"coinconv/internal/clients/coinmarketcap"
)

type Service struct {
	client *http.Client
}

func NewService() *Service {
	return &Service{client: &http.Client{}}
}

func (s *Service) PriceConversion(amount float64, from string, to string) (float64, error) {
	return coinmarketcap.NewService(s.client).PriceConversionV2(amount, from, to)
}
