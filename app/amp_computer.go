package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Nick212/cripto-currency-monitor/models/foxbit"
)

func (a *App) GetTricker(token string) (*foxbit.ResponseTricker, error) {
	computer := foxbit.ResponseTricker{}

	url := "https://api.amp.cisco.com/v1/computers"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &computer, err
	}

	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return &computer, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &computer, err
	}

	err = json.Unmarshal(body, &computer)

	return &computer, err
}
