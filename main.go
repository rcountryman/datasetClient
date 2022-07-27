package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/rcountryman/datasetClient/stockclient"
)

func main() {
	stockSymbol := os.Args[1]
	client := stockclient.StockClient{}
	client.InitStockClient()
	rawjson := client.Get(stockSymbol)

	var data stockclient.TimeSeriesData
	err := json.Unmarshal([]byte(rawjson), &data)
	if err != nil {
		log.Fatal(err)
	}

	// TODO move to data store abstraction
	// TODO conversion to csv for python
	dt := time.Now()
	_ = ioutil.WriteFile(stockSymbol+"_"+dt.Format("2006-01-02")+"_test.json", []byte(rawjson), 0644)
}
