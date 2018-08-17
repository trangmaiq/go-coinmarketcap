package types

type CryptocurrencyListingResponse struct {
	Status StatusResponse        `json:"status"`
	Data   CryptocurrencyListing `json:"data"`
}

type CryptocurrencyInfoResponse struct {
	Status StatusResponse     `json:"stauts"`
	Data   CryptocurrencyInfo `json:"data"`
}
