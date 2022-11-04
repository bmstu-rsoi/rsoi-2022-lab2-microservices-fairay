package objects

import (
	_ "encoding/json"
)

type FilghtResponse struct {
	FlightNumber string `json:"flightNumber"`
	FromAirport  string `json:"fromAirport"`
	ToAirport    string `json:"toAirport"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
}

type PaginationResponse struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	TotalElements int              `json:"totalElements"`
	Items         []FilghtResponse `json:"items"`
}

type BalanceHistory struct {
	Date          string `json:"date"`
	BalanceDiff   string `json:"balanceDiff"`
	TicketUid     string `json:"ticketUid"`
	OperationType string `json:"operationType"`
}

type PrivilegeInfoResponse struct {
	Balance string           `json:"balance"`
	Status  string           `json:"status"`
	History []BalanceHistory `json:"history"`
}