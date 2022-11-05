package models

import (
	"privileges/objects"
	"privileges/repository"
)

type PrivilegesM struct {
	rep     repository.PrivilegesRep
	history *HistoryM

	bonus_percent float32
}

func NewPrivilegesM(rep repository.PrivilegesRep, history *HistoryM) *PrivilegesM {
	return &PrivilegesM{rep, history, 0.1}
}

func (model *PrivilegesM) Find(username string) (*objects.Privilege, []objects.PrivilegeHistory, error) {
	privilege, err := model.rep.Find(username)
	if err != nil {
		return nil, nil, err
	}
	history := model.history.Find(privilege.Id)
	return privilege, history, nil
}

func (model *PrivilegesM) AddTicket(username string, info *objects.AddTicketRequest) (*objects.AddTicketResponce, error) {
	privilege, err := model.rep.Find(username)
	if err != nil {
		return nil, err
	}

	resp := &objects.AddTicketResponce{
		PaidByMoney:   info.Price,
		PaidByBonuses: 0,
	}
	balance_diff := 0
	if info.PaidFromBalance {
		if info.Price > privilege.Balance {
			resp.PaidByBonuses = privilege.Balance
		} else {
			resp.PaidByBonuses = info.Price
		}
		resp.PaidByMoney -= resp.PaidByBonuses
		balance_diff = -resp.PaidByBonuses
		err = model.history.DebitTheAccount(privilege.Id, info.TicketUID, balance_diff)
	} else {
		balance_diff = int(float32(info.Price) * model.bonus_percent)
		err = model.history.FillInBalance(privilege.Id, info.TicketUID, balance_diff)
	}
	if err != nil {
		return nil, err
	}

	privilege.Balance += balance_diff
	if err = model.rep.Update(privilege); err != nil {
		return nil, err
	} else {
		resp.Privilege.Balance = privilege.Balance
		resp.Privilege.Status = privilege.Status
		return resp, nil
	}
}
