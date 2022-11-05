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

func (model *TicketsM) Create(flight_number string, price int, from_balance bool) *objects.TicketPurchaseResponse {
	flight, _ := model.flights.Find(flight_number)

	return &objects.TicketPurchaseResponse{
		FlightNumber: flight.FlightNumber,
		FromAirport: flight.FromAirport,
		ToAirport: flight.ToAirport,
		Date: flight.Date,
		Price: price,
	}
}
