package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/whyrans/go-coinmarketcap/types"
)

func CryptocurryencyListingProAPIv1(options *ListOptions) (*types.CryptocurrencyListing, error) {
	url, errGetURL := getURL(options)
	if errGetURL != nil {
		return nil, errGetURL
	}

	log.Println(url)

	resp, errMakeReq := makeRequest(url)
	if errMakeReq != nil {
		return nil, errMakeReq
	}

	var bodyData types.CryptocurrencyResponse
	if err := json.Unmarshal(resp, &bodyData); err != nil {
		return nil, err
	}

	return &bodyData.Data, nil
}

func getURL(options *ListOptions) (string, error) {
	var params []string

	if options.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%v", options.Sort))
	}
	if options.Start >= 1 {
		params = append(params, fmt.Sprintf("start=%v", options.Start))
	}
	if options.Limit >= 1 {
		params = append(params, fmt.Sprintf("limit=%v", options.Limit))
	}
	if options.CryptocurrencyType != "" {
		params = append(params, fmt.Sprintf("cryptocurrency_type=%v", options.CryptocurrencyType))
	}
	if options.Convert != "" {
		params = append(params, fmt.Sprintf("convert=%v", options.Convert))
	}
	if options.ApiKey != "" {
		params = append(params, fmt.Sprintf("CMC_PRO_API_KEY=%v", options.ApiKey))
	}
	if options.SortDir != "" {
		params = append(params, fmt.Sprintf("sort_dir=%v", options.SortDir))
	}

	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", PRO_API_URL, strings.Join(params, "&"))
	return url, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
