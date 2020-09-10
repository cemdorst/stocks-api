package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/cemdorst/stocks-api/aao"
)

var aao = Historicals{}
var stock = StockList{}

func FindAllStocksEndPoint(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//path := "/stocks/"
	data, err := stock.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, data)
}
func HistoricalEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := "/stocks/historicals/" + params["symbol"]
	query := "?months=" + params["months"]
	historical_data, err := aao.GetHistorical(path,query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, historical_data)
}

func VolatilityEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := "/stocks/historicals/" + params["symbol"]
	query := "?months=" + params["months"]
	volatility_data, err := aao.GetVolatility(path,query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, volatility_data)
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

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.Path("/stocks").HandlerFunc(FindAllStocksEndPoint).Methods("GET")

	r.Path("/historical/{symbol}").Queries("months", "{months}").HandlerFunc(HistoricalEndPoint).Methods("GET")
	r.Path("/historical/{symbol}").HandlerFunc(HistoricalEndPoint).Methods("GET")

	r.Path("/volatility/{symbol}").Queries("months", "{months}").HandlerFunc(VolatilityEndPoint).Methods("GET")
	r.Path("/volatility/{symbol}").HandlerFunc(VolatilityEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
