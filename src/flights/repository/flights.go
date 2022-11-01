package repository

import (
	"flights/errors"
	"flights/objects"

	"github.com/jinzhu/gorm"
)

type FlightsRep interface {
	GetAll() []objects.Flight
	Find(id int) (*objects.Flight, error)
}

type PGFlightsRep struct {
	db *gorm.DB
}

func NewPGFlightsRep(db *gorm.DB) *PGFlightsRep {
	return &PGFlightsRep{db}
}

func (rep *PGFlightsRep) GetAll() []objects.Flight {
	temp := []objects.Flight{}
	rep.db.
		Model(&objects.Flight{}).
		Preload("FromAirport").
		Preload("ToAirport").
		Find(&temp)

	return temp
}

func (rep *PGFlightsRep) Find(id int) (*objects.Flight, error) {
	temp := new(objects.Flight)
	err := rep.db.Where("id = ?", id).First(temp).Error
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		temp, err = nil, errors.RecordNotFound
	default:
		temp, err = nil, errors.UnknownError
	}

	return temp, err
}
