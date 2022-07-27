# datasetClient
Go client for getting API data for creating ML datasets

To run an API key for AlphaVantage https://www.alphavantage.co/documentation/
needs to be stored in the `STOCK_API_KEY` environment variable.  This is a stopgap for not committing the key to Git.

`export STOCK_API_KEY=ThisIsYourKey`

The executable takes one command line parameter which is simply the ticker symbol of a stock

`go run main.go GOOG`

Data is stored off into a JSON file in the format symbol_date_test.json.  For example:
AAPL_2022-07-26_test.json