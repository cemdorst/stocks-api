package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/cemdorst/stocks-api/aao"
)

var aao = Historicals{}

func HistoricalEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := "/stocks/historicals/" + params["symbol"]
	historical_data, err := aao.GetHistorical(path,"")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, historical_data)
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
	r.HandleFunc("/historical/{symbol}", HistoricalEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
