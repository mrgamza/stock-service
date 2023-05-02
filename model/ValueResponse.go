package model

type ValueResponse struct {
	Index int    `json:"index"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Now   int64  `json:"nowValue"`
}
