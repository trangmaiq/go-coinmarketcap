package coinmarketcap

import (
	"testing"
)

var optionsCryptocurrencyListing = ListOptions{
	Sort:               "market_cap",
	Start:              1,
	Limit:              100,
	CryptocurrencyType: "tokens",
	Convert:            "USD",
	ApiKey:             "YOUR_API_KEY",
}

func TestCryptocurryencyListingProAPIv1(t *testing.T) {
	type args struct {
		options *ListOptions
	}
	tests := []struct {
		name string
		args args
		// want    *types.CryptocurrencyListingResponse
		wantErr bool
	}{
		{"Token Listing 1-100", args{&optionsCryptocurrencyListing}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CryptocurryencyListingProAPIv1(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurryencyListingProAPIv1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CryptocurryencyListingProAPIv1() = %v, want %v", got, tt.want)
			// }
		})
	}
}
