package models

import (
	"privileges/objects"
	"privileges/repository"
)

type PrivilegesM struct {
	rep repository.PrivilegesRep
}

func NewPrivilegesM(rep repository.PrivilegesRep) *PrivilegesM {
	return &PrivilegesM{rep}
}

func (model *PrivilegesM) Find(username string) *objects.Privilege {
	privilege, _ := model.rep.Find(username)
	return privilege
}
