package objects

import (
	_ "encoding/json"
)

type TicketPurchaseRequest struct {
	FlightNumber    string `json:"flightNumber"`
	Price           int    `json:"price"`
	PaidFromBalance bool   `json:"paidFromBalance"`
}

type TicketPurchaseResponse struct {
	TicketUid     string             `json:"ticketUid"`
	FlightNumber  string             `json:"flightNumber"`
	FromAirport   string             `json:"fromAirport"`
	ToAirport     string             `json:"toAirport"`
	Date          string             `json:"date"`
	Price         int                `json:"price"`
	PaidByMoney   int                `json:"paidByMoney"`
	PaidByBonuses int                `json:"paidByBonuses"`
	Privilege     PrivilegeShortInfo `json:"privilege"`
}

func NewTicketPurchaseResponse(flight *FlightResponse, ticket *TicketCreateResponse) *TicketPurchaseResponse {
	return &TicketPurchaseResponse{
		TicketUid: ticket.TicketUid,
		FlightNumber: flight.FlightNumber,
		FromAirport:  flight.FromAirport,
		ToAirport:    flight.ToAirport,
		Date:         flight.Date,
		Price: 		  flight.Price,
	}
}

type TicketCreateRequest struct {
	FlightNumber  string  `json:"flightNumber" gorm:"not null"`
	Price         int     `json:"price" gorm:"not null"`
}

type TicketCreateResponse struct {
	TicketUid  	  string  `json:"ticketUid"`
	Username	  string  `json:"username"`
	FlightNumber  string  `json:"flightNumber"`
	Price         int     `json:"price"`
	Status        string  `json:"status"`
}
