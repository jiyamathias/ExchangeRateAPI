package main

import (
	"fmt"
	"net/http"

	"github.com/jiyamathias/exchangerateapi"
)

func main() {
	apiKey := ""
	client := exchangerateapi.New(*http.DefaultClient, apiKey)

	sd, err := client.Standard("USD")
	if err != nil {
		fmt.Println("Error => ", err)
	}

	fmt.Println("Response ::: ", sd)

	pr, err := client.PairConversion("NGN", "USD")
	if err != nil {
		fmt.Println("Error => ", err)
	}

	fmt.Println("Response ::: ", pr)
}
