package models

import (
	"flights/errors"
	"flights/objects"
	"flights/repository"
)

type FlightsM struct {
	rep repository.FlightsRep
}

func NewFlightsM(rep repository.FlightsRep) *FlightsM {
	return &FlightsM{rep}
}

func (model *FlightsM) GetAll() []objects.Flight {
	return model.rep.GetAll()
}

func (model *FlightsM) Find(id int) (*objects.Flight, error) {
	person, err := model.rep.Find(id)
	if err != nil {
		return nil, errors.RecordNotFound
	} else {
		return person, nil
	}
}
