package coinmarketcap

import (
	"context"
	"fmt"
	"strings"
)

const (
	USD = iota
)

type CurrencyCode int

func (code CurrencyCode) String() string {
	return [...]string{
		"United States dollar",
	}[code]
}

type CryptocurrencyService service

type CryptocurrencyResponse struct {
	Status *Status           `json:"status" bson:"status"`
	Data   []*Cryptocurrency `json:"data"`
}

type Status struct {
	Timestamp    string `json:"timestamp" bson:"timestamp"`
	ErrorCode    int    `json:"error_code" bson:"error_code"`
	ErrorMessage string `json:"error_message" bson:"error_message"`
	Elapsed      int    `json:"elapsed" bson:"elapsed"`
	CreditCount  int    `json:"credit_count" bson:"credit_count"`
}

type Cryptocurrency struct {
	ID                  *int         `json:"id,omitempty" bson:"id,omitempty"`
	Name                *string      `json:"name,omitempty" bson:"name,omitempty"`
	Symbol              *string      `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Category            *string      `json:"category,omitempty" bson:"category,omitempty"`
	Slug                *string      `json:"slug,omitempty" bson:"slug,omitempty"`
	Logo                *string      `json:"logo,omitempty" bson:"logo,omitempty"`
	Tags                *[]string    `json:"tags,omitempty" bson:"tags,omitempty"`
	CmcRank             *int         `json:"cmc_rank,omitempty" bson:"cmc_rank,omitempty"`
	NumMarketPairs      *int         `json:"num_market_pairs,omitempty" bson:"num_market_pairs,omitempty"`
	MarketPairs         []MarketPair `json:"market_pairs,omitempty" bson:"market_pairs,omitempty"`
	CirculatingSupply   *float64     `json:"circulating_supply,omitempty" bson:"circulating_supply,omitempty"`
	TotalSupply         *float64     `json:"total_supply,omitempty" bson:"total_supply,omitempty"`
	MaxSupply           *int         `json:"max_supply,omitempty" bson:"max_supply,omitempty"`
	LastUpdated         *string      `json:"last_updated,omitempty" bson:"last_updated,omitempty"`
	DateAdded           *string      `json:"date_added,omitempty" bson:"date_added,omitempty"`
	Quote               *Quote       `json:"quote,omitempty" bson:"quote,omitempty"`
	URLs                *URL         `json:"urls,omitempty" bson:"urls,omitempty"`
	IsActive            *bool        `json:"is_active,omitempty" bson:"is_active,omitempty"`
	FirstHistoricalData *string      `json:"first_historical_data,omitempty" bson:"first_historical_data,omitempty"`
	LastHistoricalData  *string      `json:"last_historical_data,omitempty" bson:"last_historical_data,omitempty"`
}

type MarketPair struct {
	Exchange        *Exchange     `json:"exchange,omitempty" bson:"exchange,omitempty"`
	MarketPairKey   *string       `json:"market_pair,omitempty" bson:"market_pair,omitempty"`
	MarketPairBase  *MarketPairBQ `json:"market_pair_base,omitempty" bson:"market_pair_base,omitempty"`
	MarketPairQuote *MarketPairBQ `json:"market_pair_quote,omitempty" bson:"market_pair_quote,omitempty"`
}

type Exchange struct {
	ID     *int    `json:"id,omitempty"`
	Name   *string `json:"name,omitempty" bson:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty" bson:"symbol,omitempty"`
}

type MarketPairBQ struct {
	CurrencyID     *int    `json:"currency_id,omitempty" bson:"currency_id,omitempty"`
	CurrencySymbol *string `json:"currency_symbol,omitempty" bson:"currency_symbol,omitempty"`
	CurrencyType   *string `json:"currency_type,omitempty" bson:"currency_type,omitempty"`
}

type Quote struct {
	USD *Currency `json:"usd,omitempty" bson:"usd,omitempty"`

	//Market Pairs Quote
	ExchangeReported *ExchangeReported `json:"exchange_reported,omitempty" bson:"exchange_reported,omitempty"`
}

type ExchangeReported struct {
	Price          *float64 `json:"price,omitempty" bson:"price,omitempty"`
	Volume24hBase  *float64 `json:"volume_24h_base,omitempty" bson:"volume_24h_base,omitempty"`
	Volume24hQuote *float64 `json:"volume_24h_quote,omitempty" bson:"volume_24h_quote,omitempty"`
	LastUpdated    *string  `json:"last_updated,omitempty" bson:"last_updated,omitempty"`
}

type Currency struct {
	Price            *float64 `json:"price,omitempty" bson:"price,omitempty"`
	Volume24h        *float64 `json:"volume_24h,omitempty" bson:"volume_24h,omitempty"`
	PercentChange1h  *float64 `json:"percent_change_1h,omitempty" bson:"percent_change_1h,omitempty"`
	PercentChange24h *float64 `json:"percent_change_24h,omitempty" bson:"percent_change_24h,omitempty"`
	PercentChange7d  *float64 `json:"percent_change_7d,omitempty" bson:"percent_change_7d,omitempty"`
	MarketCap        *float64 `json:"market_cap,omitempty" bson:"market_cap,omitempty"`
}

type URL struct {
	Website      *[]string `json:"website,omitempty" bson:"website,omitempty"`
	Explorer     *[]string `json:"explorer,omitempty" bson:"explorer,omitempty"`
	SourceCode   *[]string `json:"source_code,omitempty" bson:"source_code,omitempty"`
	MessageBoard *[]string `json:"message_board,omitempty" bson:"message_board,omitempty"`
	Chat         *[]string `json:"chat,omitempty" bson:"chat,omitempty"`
	Announcement *[]string `json:"announcement,omitempty" bson:"announcement,omitempty"`
	Reddit       *[]string `json:"reddit,omitempty" bson:"reddit,omitempty"`
	Twitter      *[]string `json:"twitter,omitempty" bson:"twitter,omitempty"`
}

func (opt *ListOptions) addParamsOpts(u string) string {
	var params []string

	if opt.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%v", opt.Sort))
	}
	if opt.Start >= 1 {
		params = append(params, fmt.Sprintf("start=%v", opt.Start))
	}
	if opt.Limit >= 1 {
		params = append(params, fmt.Sprintf("limit=%v", opt.Limit))
	}
	if opt.CryptocurrencyType != "" {
		params = append(params, fmt.Sprintf("cryptocurrency_type=%v", opt.CryptocurrencyType))
	}
	if opt.Convert != "" {
		params = append(params, fmt.Sprintf("convert=%v", opt.Convert))
	}

	if opt.ApiKey != "" {
		params = append(params, fmt.Sprintf("CMC_PRO_API_KEY=%v", opt.ApiKey))
	}
	if opt.SortDir != "" {
		params = append(params, fmt.Sprintf("sort_dir=%v", opt.SortDir))
	}

	url := fmt.Sprintf("%s?%s", u, strings.Join(params, "&"))
	return url
}

func (cs *CryptocurrencyService) List(ctx context.Context, endPointPath string, opt *ListOptions) ([]*Cryptocurrency, *Response, error) {

	var u string
	if endPointPath != "" {
		u = fmt.Sprintf("cryptocurrency/listings/%s", endPointPath)
	} else {
		return nil, nil, ErrMissingEndpointPath
	}

	u = opt.addParamsOpts(u)

	req, errNewReq := cs.client.NewRequest("GET", u, nil)
	if errNewReq != nil {
		return nil, nil, errNewReq
	}

	cryptocurrency := &CryptocurrencyResponse{}
	resp, errResp := cs.client.Do(ctx, req, &cryptocurrency)
	if errResp != nil {
		return nil, nil, errResp
	}

	return cryptocurrency.Data, resp, nil
}
