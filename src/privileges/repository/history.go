package repository

import (
	"privileges/objects"

	"github.com/jinzhu/gorm"
)

type HistoryRep interface {
	Find(privilege_id int) ([]objects.PrivilegeHistory, error)
}

type PGHistoryRep struct {
	db *gorm.DB
}

func NewPGHistoryRep(db *gorm.DB) *PGHistoryRep {
	return &PGHistoryRep{db}
}

func (rep *PGHistoryRep) Find(privilege_id int) ([]objects.PrivilegeHistory, error) {
	temp := []objects.PrivilegeHistory{}
	err := rep.db.Where(objects.PrivilegeHistory{PrivilegeID: privilege_id}).Find(&temp).Error
	return temp, err
}
