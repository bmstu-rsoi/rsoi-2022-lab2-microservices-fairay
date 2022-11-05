package controllers

import (
	"gateway/controllers/responses"
	"gateway/models"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type filghtCtrl struct {
	flights *models.FlightsM
}

func InitFlights(r *mux.Router, flights *models.FlightsM) {
	ctrl := &filghtCtrl{flights}
	r.HandleFunc("/flights", ctrl.fetch).Methods("GET")
}

func (ctrl *filghtCtrl) fetch(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	page_size, _ := strconv.Atoi(queryParams.Get("size"))
	data := ctrl.flights.Fetch(page, page_size)
	responses.JsonSuccess(w, data)
}
