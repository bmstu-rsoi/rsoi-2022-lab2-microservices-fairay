package controllers

import (
	"fmt"
	"gateway/controllers/responses"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"

	"encoding/json"
	"net/http"

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

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/flights", utils.Config.FlightsEndpoint), nil)
	q := req.URL.Query()
	q.Add("page", queryParams.Get("page"))
	q.Add("size", queryParams.Get("size"))
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.PaginationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)

	responses.JsonSuccess(w, data)
}
