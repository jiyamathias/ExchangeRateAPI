package exchangerateapi

import (
	"fmt"
	"net/http"
)

// Our Standard API endpoint is the easiest and fastest way to access our exchange rate data. It's available on all plans including Free.
// Choose your base currency and the Standard endpoint will simply return the conversion rates from your base currency code to all the others we support in an easy to parse JSON format.
// We use ISO 4217 Three Letter Currency Codes - e.g. USD for US Dollars, EUR for Euro etc. Here are the :https://www.exchangerate-api.com/docs/supported-currencies
func (c *Client) Standard(currency string) (*StandardResponse, error) {
	url := fmt.Sprintf("%s/%s/latest/%s", c.baseUrl, c.apiKey, currency)

	var response StandardResponse
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Our Pair Conversion API endpoint is useful for applications where you just want to convert between two specific currencies and aren't interested in any others.
// You'll need to choose your base currency code and the target currency you'd like to get the exchange rate for. You'll then get a minimal JSON response back from our API with exactly the data you need. It's a much smaller response than the one you'd get from our Standard endpoint - quick & easy to develop with.
// We use ISO 4217 Three Letter Currency Codes - e.g. USD for US Dollars, EUR for Euro etc. Here are the :https://www.exchangerate-api.com/docs/supported-currencies
func (c *Client) PairConversion(fromCurrency, toCurrency string) (*PairConversionResponse, error) {
	url := fmt.Sprintf("%s/%s/pair/%s/%s", c.baseUrl, c.apiKey, fromCurrency, toCurrency)

	var response PairConversionResponse
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Our Enriched API endpoint offers the most comprehensive data on each currency you're interested in converting to. It's only available to users on our Business or Volume plans.
// You'll need to choose your base currency code and the target currency you'd like to get enriched data for. The API's JSON response will include the country name, country code & flag image URL for the target currency, as well as the currency name and its display symbol.
// We use ISO 4217 Three Letter Currency Codes - e.g. USD for US Dollars, EUR for Euro etc. Here are the :https://www.exchangerate-api.com/docs/supported-currencies
func (c *Client) EnrichedData(fromCurrency, toCurrency string) (*EnrichedDataResponse, error) {
	url := fmt.Sprintf("%s/%s/enriched/%s/%s", c.baseUrl, c.apiKey, fromCurrency, toCurrency)

	var response EnrichedDataResponse
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Our Historical Exchange Rates API endpoint lets you fetch all the exchange rate data we have available for a specific day in the past. It's only available to users on our Pro, Business or Volume plans.
// We use ISO 4217 Three Letter Currency Codes - e.g. USD for US Dollars, EUR for Euro etc. Here are the :https://www.exchangerate-api.com/docs/supported-currencies
// Year: 2015, Month: 11, Day: 3
func (c *Client) HistoricalDataBasicFormat(currency string, year, month, day int) (*HistoricalDataBasicFormatResponse, error) {
	url := fmt.Sprintf("%s/%s/history/%s/%d/%d/%d", c.baseUrl, c.apiKey, currency, year, month, day)

	var response HistoricalDataBasicFormatResponse
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Our Historical Exchange Rates API endpoint lets you fetch all the exchange rate data we have available for a specific day in the past. It's only available to users on our Pro, Business or Volume plans.
// We use ISO 4217 Three Letter Currency Codes - e.g. USD for US Dollars, EUR for Euro etc. Here are the :https://www.exchangerate-api.com/docs/supported-currencies
// Year: 2015, Month: 11, Day: 3, Amount: 1.4, 1...
func (c *Client) HistoricalDataSpecificFormat(currency string, year, month, day int, amount float64) (*HistoricalDataSpecificFormat, error) {
	url := fmt.Sprintf("%s/%s/history/%s/%d/%d/%d/%f", c.baseUrl, c.apiKey, currency, year, month, day, amount)

	var response HistoricalDataSpecificFormat
	if err := c.newRequest(http.MethodGet, url, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
