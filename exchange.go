package exchangerateapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey  string
	baseUrl string
	http    *http.Client
}

func New(h http.Client, apikey string) Client {
	return Client{
		apiKey:  apikey,
		baseUrl: "https://v6.exchangerate-api.com/v6",
		http:    &h,
	}
}

// newRequest makes a http request to the zeptomail server and decodes the server response into the reqBody parameter passed into the newRequest method
func (c *Client) newRequest(method, reqURL string, reqBody, resp interface{}) error {
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("http client ::: unable to marshal request struct => %v", err)
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return fmt.Errorf("http client ::: unable to create request body => %v", err)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("http client ::: client failed to execute request => %v", err)
	}
	defer res.Body.Close()

	bb, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("http client ::: client failed to read file => %v", err)
	}

	fmt.Println("Response body:", string(bb))

	if err := json.Unmarshal(bb, &resp); err != nil {
		return fmt.Errorf("http client ::: unmarshalling error => %v", err)
	}

	return nil
}
