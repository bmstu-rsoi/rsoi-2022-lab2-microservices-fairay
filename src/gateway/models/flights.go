package models

import (
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type FlightsM struct {
	client *http.Client
}

func NewFlightsM(client *http.Client) *FlightsM {
	return &FlightsM{client: client}
}

func (model *FlightsM) Fetch(page int, page_size int) *objects.PaginationResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/flights", utils.Config.FlightsEndpoint), nil)
	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("size", fmt.Sprintf("%d", page_size))
	req.URL.RawQuery = q.Encode()
	if model.client == nil {
		panic("client: nil\n")
	}

	resp, err := model.client.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.PaginationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}

func (model *FlightsM) Find(flight_number string) (*objects.FilghtResponse, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/flights/%s", utils.Config.FlightsEndpoint, flight_number), nil)
	resp, err := model.client.Do(req)
	if err != nil {
		return nil, err
	} else {
		data := &objects.FilghtResponse{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}
