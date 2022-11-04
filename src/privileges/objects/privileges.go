package objects

import "fmt"

type Privilege struct {
	Id       int    `json:"id" gorm:"primary_key;index"`
	Username string `json:"username" gorm:"not null;unique"`
	Status   string `json:"status" gorm:"not null" sql:"DEFAULT:'BRONZE'"`
	Balance  int    `json:"balance"`
}

func (Privilege) TableName() string {
	return "privilege"
}

type PrivilegeShortInfo struct {
	Balance string `json:"balance"`
	Status  string `json:"status"`
}

type PrivilegeInfoResponse struct {
	Balance string           `json:"balance"`
	Status  string           `json:"status"`
	History []BalanceHistory `json:"history"`
}

func ToPrivilegeInfoResponse(privilege *Privilege, history []PrivilegeHistory) *PrivilegeInfoResponse {
	balance_history := make([]BalanceHistory, len(history))
	for k, v := range history {
		balance_history[k] = *v.ToBalanceHistory()
	}

	return &PrivilegeInfoResponse{
		fmt.Sprintf("%d", privilege.Balance),
		privilege.Status,
		balance_history,
	}
}

type TicketResponse struct {
	TicketUid    string `json:"ticketUid"`
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
}

type UserInfoResponse struct {
	Tickets   []TicketResponse   `json:"tickets"`
	Privilege PrivilegeShortInfo `json:"privilege"`
}
