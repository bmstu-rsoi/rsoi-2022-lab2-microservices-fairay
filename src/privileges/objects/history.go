package objects

import "fmt"

type PrivilegeHistory struct {
	Id            int       `json:"id" gorm:"primary_key;index"`
	Privilege     Privilege `json:"privilege" gorm:"foreignKey:PrivilegeID"`
	PrivilegeID   int       `gorm:"index"`
	TicketUID     uint      `json:"ticketUID" gorm:"not null"`
	Datetime      string    `json:"datetime" gorm:"not null"`
	BalanceDiff   int       `json:"balanceDiff" gorm:"not null"`
	OperationType string    `json:"operationType" gorm:"not null"`
}

func (PrivilegeHistory) TableName() string {
	return "privilege_history"
}

type BalanceHistory struct {
	Date          string `json:"date"`
	BalanceDiff   string `json:"balanceDiff"`
	TicketUid     string `json:"ticketUid"`
	OperationType string `json:"operationType"`
}

func (history *PrivilegeHistory) ToBalanceHistory() *BalanceHistory {
	return &BalanceHistory{
		history.Datetime,
		fmt.Sprintf("%d", history.BalanceDiff),
		fmt.Sprintf("%d", history.TicketUID),
		history.OperationType,
	}
}
