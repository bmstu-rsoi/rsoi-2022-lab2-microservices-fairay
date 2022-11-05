package models

import (
	"tickets/errors"
	"tickets/objects"
	"tickets/repository"
)

type TicketsM struct {
	rep repository.TicketsRep
}

func NewTicketsM(rep repository.TicketsRep) *TicketsM {
	return &TicketsM{rep}
}

func (model *TicketsM) Create(user_name string, flight_number string, price int) (*objects.Ticket, error) {
	ticket := &objects.Ticket{
		Username:     user_name,
		FlightNumber: flight_number,
		Price:        price,
	}
	err := model.rep.Create(ticket)
	return ticket, err
}

func (model *TicketsM) Find(ticket_uid string) (*objects.Ticket, error) {
	person, err := model.rep.Find(ticket_uid)
	if err != nil {
		return nil, errors.RecordNotFound
	} else {
		return person, nil
	}
}
