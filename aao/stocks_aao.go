package aao

import (
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

func (v *Historicals) CalculateVariation() (*Historicals) {

	var last float64
	for i,value := range v.Data {
		if i == 0 {
			last = value.Close
			continue
		}
		diff := value.Close - last
		v.Variation = append(v.Variation,diff/last*100)
		last = value.Close
	}
	stats.StandardDeviationPopulation(v.Variation)

        return v
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
        json.Unmarshal(responseData, &responseObject)
	responseObject.CalculateVariation()

        return responseObject, err
}
