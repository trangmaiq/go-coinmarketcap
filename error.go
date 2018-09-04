package coinmarketcap

import "errors"

var (
	ErrMissingEndpointPath = errors.New("Endpoint path must not be blank.")
)
