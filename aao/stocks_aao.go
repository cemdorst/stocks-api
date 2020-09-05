package aao


import (
	"log"
	. "github.com/cemdorst/stocks-api/config"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var config = Config{}

type Historicals struct {
	Symbol string `json:"symbol"`
	Data []StocksAAO `json:"historicals"`
}


type StocksAAO struct {
	Close	float64 `json:"close"`
        Date	string `json:"date"`
        High	float64 `json:"high"`
        Low	float64 `json:"low"`
        Open	float64 `json:"open"`
        Volume	float64 `json:"volume"`
}

func (m *Historicals) GetHistorical(path,query string) (Historicals, error) {
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

