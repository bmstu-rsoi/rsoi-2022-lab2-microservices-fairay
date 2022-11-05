package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type TicketsM struct {
	client *http.Client

	flights    *FlightsM
	privileges *PrivilegesM
}

func NewTicketsM(client *http.Client, flights *FlightsM) *TicketsM {
	return &TicketsM{
		client:  client,
		flights: flights,
	}
}

func (model *TicketsM) create(flight_number string, price int) (*objects.TicketCreateResponse, error) {
	req_body, _ := json.Marshal(&objects.TicketCreateRequest{FlightNumber: flight_number, Price: price})
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/tickets", utils.Config.TicketsEndpoint), bytes.NewBuffer(req_body))

	if resp, err := model.client.Do(req); err != nil {
		return nil, err
	} else {
		data := &objects.TicketCreateResponse{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *TicketsM) Create(flight_number string, username string, price int, from_balance bool) (*objects.TicketPurchaseResponse, error) {
	flight, err := model.flights.Find(flight_number)
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	ticket, err := model.create(flight_number, price)
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	privilege, err := model.privileges.AddTicket(username, &objects.AddHistoryRequest{
		TicketUID:       ticket.TicketUid,
		Price:           flight.Price,
		PaidFromBalance: from_balance,
	})
	if err != nil {
		utils.Logger.Println(err.Error())
		return nil, err
	}

	return objects.NewTicketPurchaseResponse(flight, ticket, privilege), nil
}
