package aao

import (
	"math"
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
	Variation  []float64
	Volatility []float64
}

type StockList struct {
	*Stocks
}

func (s *StockList) FindAll() (StockList, error) {
	config.Read()
	response, err := http.Get(config.APIbase + "/stocks")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
                log.Fatal(err)
        }

        var responseObject StockList
        json.Unmarshal(responseData, &responseObject)

        return responseObject, err
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

func stdvdays(days int, variation []float64) ([]float64) {
	var stdv []float64
	var window []float64
	for range variation {
		if len(variation) > days {
			window = variation[0:days]
			stdvitem, _ := stats.StandardDeviationPopulation(window)
			stdvitem = stdvitem * math.Sqrt(252)
			stdv = append(stdv, stdvitem)
			variation = variation[1:]
		} else {
			break
		}
	}
	return stdv
}


func (v *Historicals) CalculateVolatility() (*Historicals) {

	var last float64
	for i,value := range v.Data {
		signal := 1.0
		if i == 0 {
			last = value.Close
			continue
		}
		diff := value.Close - last
		if diff < 0 {
			diff = diff * -1.0
			signal = -1.0
			v.Variation = append(v.Variation, signal * math.Log(diff))
			last = value.Close
			continue
		}
		if diff == 0 {
			v.Variation = append(v.Variation, 0.0)
			last = value.Close
			continue
		}
		v.Variation = append(v.Variation, signal * math.Log(diff))
		last = value.Close
	}

	v.Volatility = stdvdays(10, v.Variation)

        return v
}

func (v *Historicals) GetVolatility(path,query string) (Historicals, error) {
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
	responseObject.CalculateVolatility()

        return responseObject, err
}
