package model

// "marketSum" : 시가 총액. 백만단위
// "risefall" : 1 상한, 2 상승, 3 보합?, 4 하한, 5 하락
// "diff" : 전일대비 가격 차이
// "rate" : 상승율
// "high" : 고가
// "low" : 저가
// "quant" : 거래량
// "amount" : 거래대금
// "per" : PER
// "eps" : EPS
// "pbr" : PBR
// "now" : 현재가

type StockCode struct {
	MarketSum string  `json:"marketSum"`
	RiseFall  int64   `json:"risefall"`
	Diff      float64 `json:"diff"`
	Rate      float64 `json:"rate"`
	High      int64   `json:"high"`
	Low       int64   `json:"low"`
	Quant     int64   `json:"quant"`
	Amount    int64   `json:"amount"`
	PER       float64 `json:"per"`
	EPS       float64 `json:"eps"`
	PBR       float64 `json:"pbr"`
	Now       int64   `json:"now"`
}
