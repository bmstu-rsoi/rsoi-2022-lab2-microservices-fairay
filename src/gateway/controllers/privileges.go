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

type privilegeCtrl struct {
}

func InitPrivileges(r *mux.Router) {
	ctrl := &privilegeCtrl{}
	r.HandleFunc("/privilege", ctrl.get).Methods("GET")
}

func (ctrl *privilegeCtrl) get(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/privilege", utils.Config.PrivilegesEndpoint), nil)
	req.Header.Add("X-User-Name", r.Header.Get("X-User-Name"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.PrivilegeInfoResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)

	responses.JsonSuccess(w, data)
}
