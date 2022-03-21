package coinmarketcap_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"coinconv/internal/clients/coinmarketcap"

	"github.com/stretchr/testify/require"
)

func TestPriceConversionV2(t *testing.T) {
	mClient := &mockClient{
		do: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString(`
{
  "status": {
  },
  "data": [
    {
      "symbol": "BTC",
      "name": "Bitcoin",
      "amount": 1,
      "quote": {
        "USD": {
          "price": 41303.43294728996,
        }
      }
    }
  ]
}
`)),
			}, nil
		},
	}
	svc := coinmarketcap.NewService(mClient)
	res, err := svc.PriceConversionV2(1, "BTC", "USD")

	require.NoError(t, err)
	require.Equal(t, 41303.43294728996, res)
}

type mockClient struct {
	do func(req *http.Request) (*http.Response, error)
}

func (s *mockClient) Do(req *http.Request) (*http.Response, error) {
	return s.do(req)
}
