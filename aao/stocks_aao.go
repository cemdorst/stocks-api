package aao

import (
	"fmt"
	"log"
	. "github.com/cemdorst/stocks-api/config"
	. "github.com/cemdorst/stocks-api/models"
	"github.com/montanaflynn/stats"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var config = Config{}

type Historicals struct {
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

func (v *Historicals) CalculateVolatility(path,query string) (Historicals, error) {
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
	var last float64
        json.Unmarshal(responseData, &responseObject)
	for i,value := range responseObject.Data {
		if i == 0 {
			last = value.Close
			continue
		}
		responseObject.Variation = append(responseObject.Variation,value.Close/last)
		last = value.Close
	}
	a, _ := stats.StandardDeviationPopulation(responseObject.Variation)
	fmt.Println(a)

        return responseObject, err
}
