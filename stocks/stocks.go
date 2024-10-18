package stocks

type Stock struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func GetStocks() []Stock {
	stocks := []Stock{
		{
			Symbol: "AAPL",
			Name:   "Apple Inc.",
			Price:  125,
		},
		{
			Symbol: "GOOGL",
			Name:   "Alphabet Inc.",
			Price:  1750,
		},
		{
			Symbol: "AMZN",
			Name:   "Amazon.com Inc.",
			Price:  3000,
		},
		{
			Symbol: "MSFT",
			Name:   "Microsoft Corporation",
			Price:  200,
		},
	}
	return stocks
}
