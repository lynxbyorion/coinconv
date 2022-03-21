package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	client httpClient
}

func NewService(client httpClient) *Service {
	return &Service{client: client}
}

type PriceConversionV2Response struct {
	Status struct {
		ErrorCode int `json:"error_code"`
	} `json:"status"`
	Data []struct {
		Symbol string `json:"symbol"`
		Quote  map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func (s *Service) PriceConversionV2(amount float64, symbol string, convert string) (float64, error) {
	q := url.Values{}
	q.Add("amount", fmt.Sprintf("%.f", amount))
	q.Add("symbol", symbol)
	q.Add("convert", convert)

	body, err := s.getRequest("v2/tools/price-conversion", q)
	if err != nil {
		return 0, err
	}

	var resp PriceConversionV2Response
	if err := json.Unmarshal(body, &resp); err != nil {
		return 0, err
	}

	if resp.Status.ErrorCode != 0 {
		return 0, fmt.Errorf("something wrong on coinmarketcap server")
	}

	if len(resp.Data) == 0 {
		return 0, fmt.Errorf("unknown response")
	}

	convCurr, ok := resp.Data[0].Quote[convert]
	if !ok {
		return 0, fmt.Errorf("unknown response")
	}

	return convCurr.Price, nil
}

func (s *Service) getRequest(path string, values url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/"+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "36bb1f38-e611-4bc0-b42a-144cce34bcfb")
	req.URL.RawQuery = values.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed sending request: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read body: %w", err)
	}

	return respBody, nil
}
