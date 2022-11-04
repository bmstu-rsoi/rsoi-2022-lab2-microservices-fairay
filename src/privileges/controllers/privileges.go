package controllers

import (
	"privileges/controllers/responses"
	"privileges/models"
	"privileges/objects"

	"net/http"

	"github.com/gorilla/mux"
)

type privilegesCtrl struct {
	privileges *models.PrivilegesM
	history    *models.HistoryM
}

func InitPrivileges(r *mux.Router, privileges *models.PrivilegesM, history *models.HistoryM) {
	ctrl := &privilegesCtrl{privileges, history}
	r.HandleFunc("/privilege", ctrl.get).Methods("GET")
}

func (ctrl *privilegesCtrl) get(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Name")

	privilege := ctrl.privileges.Find(username)
	history := ctrl.history.Find(privilege.Id)

	data := objects.ToPrivilegeInfoResponse(privilege, history)
	responses.JsonSuccess(w, data)
}
