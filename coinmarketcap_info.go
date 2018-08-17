package coinmarketcap

import (
	"encoding/json"

	"github.com/whyrans/go-coinmarketcap/types"
)

func CryptocurryencyInfoProAPIv1(options *InfoOptions) (*types.CryptocurrencyInfoResponse, error) {
	url, errGetURL := getURL(&Option{InfoOptions: *options})

	if errGetURL != nil {
		return nil, errGetURL
	}

	resp, errMakeReq := makeRequest(url)
	if errMakeReq != nil {
		return nil, errMakeReq
	}

	var bodyData types.CryptocurrencyInfoResponse
	if err := json.Unmarshal(resp, &bodyData); err != nil {
		return nil, err
	}

	return &bodyData, nil
}
