package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/cemdorst/stocks-api/config"
	. "github.com/cemdorst/stocks-api/dao"
)

var config = Config{}
var dao = StocksDAO{}

// GET list of stocks
func AllStocksEndPoint(w http.ResponseWriter, r *http.Request) {
	stocks, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, stocks)
}

func HistoricalEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	historical_data, err := dao.FindHistoricalBySymbol(params["symbol"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, historical_data)
}

func VolatilityEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data, err := dao.FindHistoricalBySymbol(params["symbol"])
	fmt.Println(data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, data)
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/stocks", AllStocksEndPoint).Methods("GET")
	r.HandleFunc("/historical/{symbol}", HistoricalEndPoint).Methods("GET")
	r.HandleFunc("/volatility/{symbol}", VolatilityEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
