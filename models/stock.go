package models

import "time"

type StockHistoricals struct {
        Symbol string     `json:"symbol"`
        Data []Historical `json:"historicals"`
}


type Historical struct {
        Close   float64    `json:"close"`
        Date    time.Time  `json:"date"`
        High    float64    `json:"high"`
        Low     float64    `json:"low"`
        Open    float64    `json:"open"`
        Volume  float64    `json:"volume"`
}
