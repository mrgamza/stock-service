package stock

import (
	"github.com/gorilla/mux"
	"net/http"
	"stock-service/route"
	service "stock-service/service/naver"
)

func Route(router *mux.Router) {
	subRouter := router.PathPrefix("/stock").
		Methods("GET").
		Subrouter()

	route.HandleQueryFunc(subRouter, "keyword", getKeyword)
	route.HandleQueryFunc(subRouter, "code", getCode)
	route.HandleQueryFunc(subRouter, "query", getQuery)
}

func getKeyword(writer http.ResponseWriter, request *http.Request) {
	keyword := route.GetPath(request, "keyword")
	response, err := service.GetKeyword(keyword)
	data, err := route.MakeInterface(response)
	route.WriteResponse(writer, request, data, err)
}

func getCode(writer http.ResponseWriter, request *http.Request) {
	code := route.GetPath(request, "code")
	response, err := service.GetCode(code)
	data, err := route.MakeInterface(response)
	route.WriteResponse(writer, request, data, err)
}

func getQuery(writer http.ResponseWriter, request *http.Request) {
	query := route.GetPath(request, "query")
	response, err := service.GetQuery(query)
	route.WriteResponse(writer, request, response, err)
}
