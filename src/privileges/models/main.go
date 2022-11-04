package models

import (
	"privileges/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Privileges *PrivilegesM
	History    *HistoryM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Privileges = NewPrivilegesM(repository.NewPGPrivilegesRep(db))
	models.History = NewHistoryM(repository.NewPGHistoryRep(db))

	return models
}
