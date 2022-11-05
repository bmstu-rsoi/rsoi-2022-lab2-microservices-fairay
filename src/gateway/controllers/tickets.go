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
	flights *models.FlightsM
}

func InitTickets(r *mux.Router, flights *models.FlightsM) {
	ctrl := &ticketsCtrl{flights: flights}
	r.HandleFunc("/tickets", ctrl.post).Methods("POST")
}

func (ctrl *ticketsCtrl) post(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.TicketPurchaseRequest)
	err := json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	if data, err := ctrl.flights.Find(req_body.FlightNumber); err != nil {
		responses.InternalError(w)
	} else {
		responses.JsonSuccess(w, data)
	}
}
