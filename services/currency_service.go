package services

import (
	"boletia/models"
	"os"

	"github.com/go-resty/resty/v2"
)

type CurrencyService struct {
	CurrencyData *models.CurrencyData
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		CurrencyData: &models.CurrencyData{},
	}
}

func (s *CurrencyService) UpdateCurrencyData() {
	apiKey := os.Getenv("APIKEY")
	apiURL := os.Getenv("API_URL")
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
			"code":   "MXN",
		}).Get(apiURL)
	if err != nil {
		return
	}

	var data models.CurrencyData
	err2 := resp.Result()

	s.CurrencyData = &data
}
