# ExchangeRateAPI
Golang sdk for https://exchangerate-api.com

# Get Started
In order to use this package, you need to first head over to https://www.exchangerate-api.com to create an account to get your `API Key`.
# Documentation
See the [Exchange Rates API Docs](https://www.exchangerate-api.com/docs/overview)

# Installation
This package can be installed using the go command below.
```sh
go get github.com/iqquee/exchangerateapi@latest
```
# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## Features
- Standard()
- PairConversion()
- EnrichedData()
- HistoricalDataBasicFormat()
- HistoricalDataSpecificFormat()
  
# Import 
```go
    "github.com/iqquee/exchangerateapi"
```
# Standard
Our Standard API endpoint is the easiest and fastest way to access our exchange rate data.
```go
    apiKey := ""
	client := exchangerateapi.New(*http.DefaultClient, apiKey)

	sd, err := client.Standard("USD")
	if err != nil {
		fmt.Println("Error => ", err)
	}

	fmt.Println("Response ::: ", sd)
```

# PairConversion
Our Pair Conversion API endpoint is useful for applications where you just want to convert between two specific currencies and aren't interested in any others.
```go
    apiKey := ""
	client := exchangerateapi.New(*http.DefaultClient, apiKey)

	pr, err := client.PairConversion("NGN", "USD")
	if err != nil {
		fmt.Println("Error => ", err)
	}

	fmt.Println("Response ::: ", pr)
```
