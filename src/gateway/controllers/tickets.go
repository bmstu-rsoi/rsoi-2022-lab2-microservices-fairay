package controllers

import (
	"gateway/controllers/responses"
	"gateway/models"
	"gateway/objects"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ticketsCtrl struct {
	tickets *models.TicketsM
}

func InitTickets(r *mux.Router, tickets *models.TicketsM) {
	ctrl := &ticketsCtrl{tickets: tickets}
	r.HandleFunc("/tickets", ctrl.post).Methods("POST")
}

func (ctrl *ticketsCtrl) post(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.TicketPurchaseRequest)
	err := json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	data, err := ctrl.tickets.Create(req_body.FlightNumber, r.Header.Get("X-User-Name"), req_body.Price, req_body.PaidFromBalance)
	if err != nil {
		responses.InternalError(w)
	} else {
		responses.JsonSuccess(w, data)
	}
}