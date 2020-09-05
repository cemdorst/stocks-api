package aao

import (
	"log"
	"fmt"

	//. "github.com/cemdorst/stocks-api/models"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	API_BASE_URL = "https://mfinance.com.br/api/v1"
)

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

//type Consume interface {
//	GetHistorical()
//}

func (m *Historicals) GetHistorical(path,query string) (Historicals, error) {
	response, err := http.Get(API_BASE_URL + path + query)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Historicals
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseData)
	fmt.Println(responseObject.Data)

	return responseObject, err
}

