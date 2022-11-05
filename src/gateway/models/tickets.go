package models

import (
	"gateway/objects"
	"net/http"
)

type TicketsM struct {
	client *http.Client
	flights *FlightsM
}

func NewTicketsM(client *http.Client, flights *FlightsM) *TicketsM {
	return &TicketsM{
		client: client,
		flights: flights,
	}
}

func (model *TicketsM) Create(flight_number string, user_name string, price int, from_balance bool) (*objects.TicketPurchaseResponse, error) {
	flight, err := model.flights.Find(flight_number)
	if err != nil {
		return nil, err
	}

	
	return objects.NewTicketPurchaseResponse(flight), nil
}
