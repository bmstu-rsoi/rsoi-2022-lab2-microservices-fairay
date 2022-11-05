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

func (model *TicketsM) Fetch() []objects.Ticket {
	return model.rep.Fetch()
}

func (model *TicketsM) Create(username string, flight_number string, price int) (*objects.Ticket, error) {
	ticket := &objects.Ticket{
		Username:     username,
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
