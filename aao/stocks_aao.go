package aao

import (
	"log"
	//"fmt"

	//. "github.com/cemdorst/stocks-api/models"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
	"io/ioutil"
)

const (
	API_BASE_URL = "http://mfinance.com.br/api/v1"
)

type StocksAAO struct {
	Symbol string  `json:"symbol"`
	Close	float64 `json:"close"`
        Date	string `json:"date"`
        High	float64 `json:"high"`
        Low	float64 `json:"low"`
        Open	float64 `json:"open"`
        Volume	float64 `json:"volume"`
}

type Consume interface {
	GetHistorical() []byte
}

func (m *StocksAAO) GetHistorical(path,query string) ([]byte) {
	response, err := http.Get(API_BASE_URL + path + query)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

