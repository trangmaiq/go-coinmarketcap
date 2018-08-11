package coinmarketcap

type ListOptions struct {
	Start              int
	Limit              int
	Convert            string
	Sort               string
	SortDir            string
	CryptocurrencyType string
	ApiKey             string
}
