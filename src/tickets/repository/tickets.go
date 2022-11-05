package repository

import (
	"tickets/errors"
	"tickets/objects"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type TicketsRep interface {
	Create(*objects.Ticket) error
	Find(ticket_uid string) (*objects.Ticket, error)
}

type PGTicketsRep struct {
	db *gorm.DB
}

func NewPGTicketsRep(db *gorm.DB) *PGTicketsRep {
	return &PGTicketsRep{db}
}

func (rep *PGTicketsRep) Create(ticket *objects.Ticket) error {
	ticket.TicketUid = uuid.New().String()
	ticket.Status = "PAID"

	return rep.db.Create(ticket).Error
}

func (rep *PGTicketsRep) Find(ticket_uid string) (*objects.Ticket, error) {
	temp := new(objects.Ticket)
	err := rep.db.
		Where(&objects.Ticket{TicketUid: ticket_uid}).
		First(temp).
		Error
	switch err {
	case nil:
		return temp, err
	case gorm.ErrRecordNotFound:
		return nil, errors.RecordNotFound
	default:
		return nil, errors.UnknownError
	}
}
