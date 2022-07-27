package main

import (
	"fmt"
	"os"

	"github.com/rcountryman/datasetClient/stockclient"
)

func main() {
	stockSymbol := os.Args[1]
	client := stockclient.StockClient{}
	client.InitStockClient()
	json := client.Get(stockSymbol)
	fmt.Println(json)
}
