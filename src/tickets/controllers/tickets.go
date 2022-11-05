package controllers

import (
	"encoding/json"
	"tickets/controllers/responses"
	"tickets/errors"
	"tickets/models"
	"tickets/objects"

	"net/http"

	"github.com/gorilla/mux"
)

type filghtCtrl struct {
	model *models.TicketsM
}

func InitTickets(r *mux.Router, model *models.TicketsM) {
	ctrl := &filghtCtrl{model}
	r.HandleFunc("/tickets", ctrl.create).Methods("POST")
	r.HandleFunc("/tickets/{ticketNumber}", ctrl.get).Methods("GET")
}

func (ctrl *filghtCtrl) create(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.CreateRequest)
	json.NewDecoder(r.Body).Decode(req_body)

	ticket, _ := ctrl.model.Create(r.Header.Get("X-User-Name"), req_body.FlightNumber, req_body.Price)
	responses.JsonSuccess(w, ticket)
}

func (ctrl *filghtCtrl) get(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	ticket_uid := urlParams["ticketUid"]

	data, err := ctrl.model.Find(ticket_uid)
	switch err {
	case nil:
		responses.JsonSuccess(w, data)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, ticket_uid)
	default:
		responses.InternalError(w)
	}
}