package model

type StockKeyword struct {
	ResultCode string             `json:"resultCode"`
	Result     StockKeywordResult `json:"result"`
}

type StockKeywordResult struct {
	Data       []StockKeywordData `json:"d"`
	TotalCount int64              `json:"totCnt"`
	Type       string             `json:"search"`
}

type StockKeywordData struct {
	Code                  string  `json:"cd"`
	Name                  string  `json:"nm"`
	High                  string  `json:"nv"`
	ComparedPreviousValue string  `json:"cv"`
	ComparedPreviousRate  string  `json:"cr"`
	RiseFall              string  `json:"rf"`
	MarketCap             float64 `json:"mks"`
	TransactionPrice      float64 `json:"aa"`
	Nation                string  `json:"nation"`
	Etf                   bool    `json:"etf"`
}
