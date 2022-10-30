package controllers

import (
	"gateway/controllers/responses"
	"gateway/objects"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type filghtCtrl struct {
}

func InitFlights(r *mux.Router) {
	ctrl := &filghtCtrl{}
	r.HandleFunc("/flights", ctrl.get).Methods("GET")
}

func (ctrl *filghtCtrl) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	page_size, _ := strconv.Atoi(queryParams.Get("size"))

	data := &objects.PaginationResponse{
		Page:          page,
		PageSize:      page_size,
		TotalElements: 1,
		Items: []objects.FilghtResponse{
			{
				FlightNumber: "AFL031",
				FromAirport:  "Санкт-Петербург Пулково",
				ToAirport:    "Москва Шереметьево",
				Date:         "2021-10-08 20:00",
				Price:        1500,
			},
		},
	}

	responses.JsonSuccess(w, data)
}
