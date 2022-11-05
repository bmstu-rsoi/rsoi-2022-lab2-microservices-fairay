package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/errors"
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

func (model *TicketsM) create(flight_number string, price int, username string) (*objects.TicketCreateResponse, error) {
	req_body, _ := json.Marshal(&objects.TicketCreateRequest{FlightNumber: flight_number, Price: price})
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/tickets", utils.Config.TicketsEndpoint), bytes.NewBuffer(req_body))
	req.Header.Add("X-User-Name", username)

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

	ticket, err := model.create(flight_number, price, username)
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

func (model *TicketsM) find(ticket_uid string) (*objects.Ticket, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/tickets/%s", utils.Config.TicketsEndpoint, ticket_uid), nil)
	resp, err := model.client.Do(req)
	if err != nil {
		return nil, err
	} else {
		data := &objects.Ticket{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *TicketsM) Find(ticket_uid string, username string) (*objects.TicketResponse, error) {
	ticket, err := model.find(ticket_uid)
	if err != nil {
		return nil, err
	} else if username != ticket.Username {
		utils.Logger.Printf("Username %s != %s", username, ticket.Username)
		return nil, errors.ForbiddenTicket
	}

	flight, err := model.flights.Find(ticket.FlightNumber)
	if err != nil {
		return nil, err
	} else {
		return objects.ToTicketResponce(ticket, flight), nil
	}
}
