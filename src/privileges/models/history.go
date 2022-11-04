package models

import (
	"privileges/objects"
	"privileges/repository"
)

type HistoryM struct {
	rep repository.HistoryRep
}

func NewHistoryM(rep repository.HistoryRep) *HistoryM {
	return &HistoryM{rep}
}

func (model *HistoryM) Find(privilege_id int) []objects.PrivilegeHistory {
	history, _ := model.rep.Find(privilege_id)
	return history
}
