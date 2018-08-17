package coinmarketcap

type Option struct {
	ListOptions
	InfoOptions
	ApiKey string
}
type ListOptions struct {
	Start              int
	Limit              int
	Convert            string
	Sort               string
	SortDir            string
	CryptocurrencyType string
}

type InfoOptions struct {
	ID     int
	Symbol string
}
