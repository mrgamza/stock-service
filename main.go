package main

import (
	"log"
	"net/http"
	"os"
	"stock-service/route"
	"stock-service/route/stock"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", route.HandleHello).Methods("GET")
	stock.Route(router)

	log.Print("Start stock-service")

	err := http.ListenAndServe(":"+port(), router)
	log.Fatal(err)
}

func port() string {
	port := os.Getenv("PORT")
	if len(strings.TrimSpace(port)) != 0 {
		return port
	} else {
		return "5000"
	}
}
