package stockclient

type TimeSeriesData struct {
	MetaData   MetaDataStruct             `json:"Meta Data"`
	TimeSeries map[string]StockDataStruct `json:"Time Series (Daily)"`
}

type MetaDataStruct struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type StockDataStruct struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}
