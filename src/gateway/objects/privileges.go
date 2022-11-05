package objects

import (
	_ "encoding/json"
)

type BalanceHistory struct {
	Date          string `json:"date"`
	BalanceDiff   string `json:"balanceDiff"`
	TicketUid     string `json:"ticketUid"`
	OperationType string `json:"operationType"`
}

type PrivilegeShortInfo struct {
	Balance 			string 			`json:"balance"`
	Status  			string 			`json:"status"`
}

type PrivilegeInfoResponse struct {
	Balance string           `json:"balance"`
	Status  string           `json:"status"`
	History []BalanceHistory `json:"history"`
}