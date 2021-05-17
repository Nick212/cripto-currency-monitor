package foxbit

type ResponseTicker struct {
	Currency      string  `json:"currency,omitempty"`
	Exchange      string  `json:"exchange,omitempty"`
	Countrycode   string  `json:"countryCode,omitempty"`
	Icon          string  `json:"icon,omitempty"`
	High          float64 `json:"high,omitempty"`
	Low           float64 `json:"low,omitempty"`
	Last          float64 `json:"last,omitempty"`
	Lastvariation float64 `json:"lastVariation,omitempty"`
	Vol           float64 `json:"vol,omitempty"`
	Buyprice      float64 `json:"buyPrice,omitempty"`
	Sellprice     float64 `json:"sellPrice,omitempty"`
	Averageprice  float64 `json:"averagePrice,omitempty"`
	Createddate   string  `json:"createdDate,omitempty"`
}
