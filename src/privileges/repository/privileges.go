package repository

import (
	"privileges/errors"
	"privileges/objects"

	"github.com/jinzhu/gorm"
)

type PrivilegesRep interface {
	Find(username string) (*objects.Privilege, error)
}

type PGPrivilegesRep struct {
	db *gorm.DB
}

func NewPGPrivilegesRep(db *gorm.DB) *PGPrivilegesRep {
	return &PGPrivilegesRep{db}
}

func (rep *PGPrivilegesRep) Find(username string) (*objects.Privilege, error) {
	temp := new(objects.Privilege)
	err := rep.db.Where(objects.Privilege{Username: username}).First(temp).Error
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
