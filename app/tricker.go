package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Nick212/cripto-currency-monitor/models/foxbit"
)

func (a *App) GetTricker() (*foxbit.ResponseTicker, error) {
	tricker := foxbit.ResponseTicker{}

	url := a.Config.HOST_API + "Ticker?exchange=Foxbit&Pair=BRLXBTC"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &tricker, err
	}

	// req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return &tricker, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &tricker, err
	}

	err = json.Unmarshal(body, &tricker)

	return &tricker, err
}
