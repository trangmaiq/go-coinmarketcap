package coinmarketcap

import (
	"testing"
)

var optionsCryptocurrencyInfo = InfoOptions{
	Symbol: "ETH",
	ApiKey: "YOUR_API_KEY",
}

func TestCryptocurryencyInfoProAPIv1(t *testing.T) {
	type args struct {
		options *InfoOptions
	}
	tests := []struct {
		name string
		args args
		// want    *types.CryptocurrencyInfoResponse
		wantErr bool
	}{
		{"Get Information of cryptocurrency by symbol", args{&optionsCryptocurrencyInfo}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CryptocurryencyInfoProAPIv1(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("CryptocurryencyInfoProAPIv1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CryptocurryencyInfoProAPIv1() = %v, want %v", got, tt.want)
			// }
		})
	}
}
