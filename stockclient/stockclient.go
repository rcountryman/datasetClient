package stockclient

import (
	"io/ioutil"
	"net/http"
	"os"
)

type StockClient struct {
	apiKey string
}

func (client *StockClient) InitStockClient() {
	client.apiKey = os.Getenv("STOCK_API_KEY")
}

func (client StockClient) Get(stock string) (stringResp string) {
	if client.apiKey == "" {
		return "Not authorized"
	}

	if stock == "" {
		return "No stock ticker on input"
	}

	resp, err := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=" + stock + "&interval=60min&apikey=" + client.apiKey)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body)
}
