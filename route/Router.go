package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/matryer/respond"
	"net/http"
	"stock-service/model"
)

func HandleHello(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "Hello, stock-service!")
}

func HandleFunc(router *mux.Router, f func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return router.HandleFunc("", f)
}

func HandleQueryFunc(router *mux.Router, query string, f func(http.ResponseWriter,
	*http.Request)) *mux.Route {
	return HandleFunc(router, f).Queries(query, "{"+query+"}")
}

func GetPath(request *http.Request, name string) string {
	vars := mux.Vars(request)
	return vars[name]
}

func MakeInterface(response *http.Response) (interface{}, error) {
	var data interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func WriteResponse(writer http.ResponseWriter, request *http.Request, response interface{}, err error) {
	if err != nil {
		writeError(writer, request, err)
	} else {
		write(writer, request, response)
	}
}

func write(writer http.ResponseWriter, request *http.Request, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	responseModel := model.CommonResponse{
		Code:    0000,
		Message: "success",
		Data:    data,
	}

	respond.With(writer, request, http.StatusOK, responseModel)
}

func writeError(writer http.ResponseWriter, request *http.Request, err error) {
	writer.Header().Set("Content-Type", "application/json")
	data := model.CommonResponse{
		Code:    1000,
		Message: err.Error(),
	}

	respond.With(writer, request, http.StatusInternalServerError, data)
}
