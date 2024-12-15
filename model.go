package exchangerateapi

type (
	// StandardResponse object
	StandardResponse struct {
		Result             string             `json:"result"`
		Documentation      string             `json:"documentation"`
		TermsOfUse         string             `json:"terms_of_use"`
		TimeLasUpdateUnix  int                `json:"time_last_update_unix"`
		TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
		TimeNextUpdateUnix int                `json:"time_next_update_unix"`
		TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
		BaseCode           string             `json:"base_code"`
		ConversionRates    map[string]float64 `json:"conversion_rates"`
		ErrorType          string             `json:"error-type"`
	}

	PairConversionResponse struct {
		Result             string  `json:"result"`
		Documentation      string  `json:"documentation"`
		TermsOfUse         string  `json:"terms_of_use"`
		TimeLasUpdateUnix  int64   `json:"time_last_update_unix"`
		TimeLastUpdateUTC  string  `json:"time_last_update_utc"`
		TimeNextUpdateUnix int64   `json:"time_next_update_unix"`
		TimeNextUpdateUTC  string  `json:"time_next_update_utc"`
		BaseCode           string  `json:"base_code"`
		TargetCode         string  `json:"target_code"`
		ConversionRate     float64 `json:"conversion_rate"`
		ErrorType          string  `json:"error-type"`
	}

	EnrichedDataResponse struct {
		Result             string  `json:"result"`
		Documentation      string  `json:"documentation"`
		TermsOfUse         string  `json:"terms_of_use"`
		TimeLasUpdateUnix  int64   `json:"time_last_update_unix"`
		TimeLastUpdateUTC  string  `json:"time_last_update_utc"`
		TimeNextUpdateUnix int64   `json:"time_next_update_unix"`
		TimeNextUpdateUTC  string  `json:"time_next_update_utc"`
		BaseCode           string  `json:"base_code"`
		TargetCode         string  `json:"target_code"`
		ConversionRate     float64 `json:"conversion_rate"`
		TargetData         struct {
			Locale            string `json:"locale"`
			TwoLetterCode     string `json:"two_letter_code"`
			CurrencyName      string `json:"currency_name"`
			CurrencyNameShort string `json:"currency_name_short"`
			DisplaySymbol     string `json:"display_symbol"`
			FlagURL           string `json:"flag_url"`
		} `json:"target_data"`
		ErrorType string `json:"error-type"`
	}

	HistoricalDataBasicFormatResponse struct {
		Result          string             `json:"result"`
		Documentation   string             `json:"documentation"`
		TermsOfUse      string             `json:"terms_of_use"`
		Year            int                `json:"year"`
		Month           int                `json:"month"`
		Day             int                `json:"day"`
		BaseCode        string             `json:"base_code"`
		ConversionRates map[string]float64 `json:"conversion_rates"`
	}

	HistoricalDataSpecificFormat struct {
		Result            string             `json:"result"`
		Documentation     string             `json:"documentation"`
		TermsOfUse        string             `json:"terms_of_use"`
		Year              int                `json:"year"`
		Month             int                `json:"month"`
		Day               int                `json:"day"`
		BaseCode          string             `json:"base_code"`
		RequestedAmount   float64            `json:"requested_amount"`
		ConversionAmounts map[string]float64 `json:"conversion_amounts"`
	}
)
