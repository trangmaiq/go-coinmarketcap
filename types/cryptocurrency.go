package types

import "time"

type Cryptocurrency struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Symbol            string              `json:"symbol"`
	Slug              string              `json:"slug"`
	CirculatingSupply float64             `json:"circulating_supply"`
	TotalSupply       float64             `json:"total_supply"`
	MaxSupply         interface{}         `json:"max_supply"`
	DateAdded         time.Time           `json:"date_added"`
	NumMarketPairs    int                 `json:"num_market_pairs"`
	CmcRank           int                 `json:"cmc_rank"`
	LastUpdated       time.Time           `json:"last_updated"`
	Quote             CryptoCurrencyQuote `json:"quote"`
}

type CryptoCurrencyQuote struct {
	USD USD `json:"USD"`
}

type USD struct {
	Price            float64   `json:"price"`
	Volume24H        float64   `json:"volume_24h"`
	PercentChange1H  float64   `json:"percent_change_1h"`
	PercentChange24H float64   `json:"percent_change_24h"`
	PercentChange7D  float64   `json:"percent_change_7d"`
	MarketCap        float64   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}

type CryptocurrencyListing []Cryptocurrency
type CryptocurrencyResponse struct {
	Status StatusResponse        `json:"status"`
	Data   CryptocurrencyListing `json:"data"`
}

type StatusResponse struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}
