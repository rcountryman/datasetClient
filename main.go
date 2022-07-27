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
	// Really don't need to unmarshal JSON right now, but keeping for when move data to CSV
	err := json.Unmarshal([]byte(rawjson), &data)
	if err != nil {
		log.Fatal(err)
	}

	// TODO conversion to csv for python
	// TODO Add either datastore or folder structure
	dt := time.Now()
	_ = ioutil.WriteFile(stockSymbol+"_"+dt.Format("2006-01-02")+"_timeseriesdata.json", []byte(rawjson), 0644)
}
