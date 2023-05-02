package naver

import (
	"encoding/json"
	"fmt"
	"net/http"
	errors "stock-service/error"
	"stock-service/helper"
	"stock-service/model"
)

func GetKeyword(keyword string) (*http.Response, error) {
	length := len(keyword)

	if length < 1 {
		return nil, &errors.NaverError{Message: "keyword is empty."}
	}

	encode := helper.Encode(keyword)

	keywordUrl := "http://m.stock.naver.com/api/json/search/searchListJson.nhn?keyword"
	url := fmt.Sprintf("%s=%s", keywordUrl, encode)
	response, err := requestGet(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetCode(code string) (*http.Response, error) {
	length := len(code)

	if length < 1 {
		return nil, &errors.NaverError{Message: "keyword is empty."}
	}

	codeUrl := "http://api.finance.naver.com/service/itemSummary.nhn?itemcode"
	url := fmt.Sprintf("%s=%s", codeUrl, code)
	response, err := requestGet(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetQuery(query string) (interface{}, error) {
	keywordResponse, err := GetKeyword(query)
	if err != nil {
		return nil, err
	}

	var data model.StockKeyword
	unmarshalErr := json.NewDecoder(keywordResponse.Body).Decode(&data)
	if unmarshalErr != nil {
		return nil, err
	}

	var result []model.ValueResponse

	for index, data := range data.Result.Data {
		codeResponse, _ := GetCode(data.Code)
		var valueData model.StockCode
		unmarshalErr := json.NewDecoder(codeResponse.Body).Decode(&valueData)
		if unmarshalErr != nil {
			value := model.ValueResponse{
				Index: index,
				Code:  data.Code,
				Name:  data.Name,
				Now:   valueData.Now,
			}

			result = append(result, value)
		}
	}

	return result, nil
}

func requestGet(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if request == nil || err != nil {
		return nil, &errors.NaverError{Message: "request GET error"}
	}

	client := &http.Client{}
	return client.Do(request)
}
