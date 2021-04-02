package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Quote struct {
	GlobalQuote struct {
		Symbol           string  `json:"01. symbol"`
		Open             float32 `json:"02. open"`
		High             float32 `json:"03. high"`
		Low              float32 `json:"04. low"`
		Price            string  `json:"05. price"`
		Volume           int32   `json:"06. volume"`
		LatestTradingDay string  `json:"07. latest trading day"`
		PreviousClose    float32 `json:"08. previous close"`
		Change           float32 `json:"09. change"`
		ChangePercent    string  `json:"10. change percent"`
	} `json:"Global Quote"`
}

func (c *Client) GetQuote(symbol string) (Quote, error) {
	var quote Quote

	parameters := url.Values{}
	parameters.Add("function", "GLOBAL_QUOTE")
	parameters.Add("symbol", symbol)
	parameters.Add("apikey", c.ApiKey)
	c.BaseUrl.RawQuery = parameters.Encode()

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		return quote, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return quote, err
	}

	err = json.NewDecoder(res.Body).Decode(&quote)
	if err != nil {
		return quote, err
	}

	return quote, nil
}