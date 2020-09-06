package aao

import (
	"log"
	. "github.com/cemdorst/stocks-api/config"
	. "github.com/cemdorst/stocks-api/models"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var config = Config{}

type Historicals StockHistoricals

type Volatility struct {
	*StockHistoricals
	Variation []float64
}

func (h *Historicals) GetHistorical(path,query string) (Historicals, error) {
	config.Read()
	response, err := http.Get(config.APIbase + path + query)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Historicals
	json.Unmarshal(responseData, &responseObject)

	return responseObject, err
}

func (v *Volatility) CalculateVolatility(path,query string) (Volatility, error) {
	config.Read()
        response, err := http.Get(config.APIbase + path + query)
        if err != nil {
                log.Fatal(err)
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
                log.Fatal(err)
        }

        var responseObject Volatility
        json.Unmarshal(responseData, &responseObject)

        return responseObject, err
}
