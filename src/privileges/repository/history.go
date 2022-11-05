package repository

import (
	"privileges/objects"
	"time"

	"github.com/jinzhu/gorm"
)

type HistoryRep interface {
	Create(entry *objects.PrivilegeHistory) error
	Find(privilege_id int) ([]objects.PrivilegeHistory, error)
}

type PGHistoryRep struct {
	db *gorm.DB
}

func NewPGHistoryRep(db *gorm.DB) *PGHistoryRep {
	return &PGHistoryRep{db}
}

func (rep *PGHistoryRep) Create(entry *objects.PrivilegeHistory) error {
	entry.Datetime = time.Now().Format(time.RFC3339)
	return rep.db.Create(entry).Error
}

func (rep *PGHistoryRep) Find(privilege_id int) ([]objects.PrivilegeHistory, error) {
	temp := []objects.PrivilegeHistory{}
	err := rep.db.Where(objects.PrivilegeHistory{PrivilegeID: privilege_id}).Find(&temp).Error
	return temp, err
}
