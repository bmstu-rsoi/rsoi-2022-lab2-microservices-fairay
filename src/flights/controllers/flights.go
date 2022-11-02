package controllers

import (
	"flights/controllers/responses"
	"flights/models"
	"flights/objects"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type filghtCtrl struct {
	model *models.FlightsM
}

func InitFlights(r *mux.Router, model *models.FlightsM) {
	ctrl := &filghtCtrl{model}
	r.HandleFunc("/flights", ctrl.getAll).Methods("GET")
}

func (ctrl *filghtCtrl) getAll(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	page_size, _ := strconv.Atoi(queryParams.Get("size"))
	items := ctrl.model.GetAll(page, page_size)

	data := &objects.PaginationResponse{
		Page:          page,
		PageSize:      page_size,
		TotalElements: len(items),
		Items:         objects.ToFilghtResponses(items),
	}

	responses.JsonSuccess(w, data)
}
