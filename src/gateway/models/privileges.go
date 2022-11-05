package models

import (
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type PrivilegesM struct {
	client *http.Client
}

func NewPrivilegesM(client *http.Client) *PrivilegesM {
	return &PrivilegesM{client: client}
}

func (model *PrivilegesM) Fetch(user_name string) *objects.PrivilegeInfoResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/privilege", utils.Config.PrivilegesEndpoint), nil)
	req.Header.Add("X-User-Name", user_name)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.PrivilegeInfoResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}
