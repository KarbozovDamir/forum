package router

import (
	"encoding/json"
	"errors"
	"net/http"

	data "github.com/KarbozovDamir/Forum/internal/data"
)

//Users - special struct for db

//Data - needs to send data to our html
type Data struct {
	CheckUser, CheckMail, CheckPsw string
}

// AjaxHandler - to check is username and mail taken
func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ErrorHandler(w, r, errors.New("page for ajax"), 2)
	} else {
		var d Data
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if d.CheckUser != "" {
			_, err := data.GetUserByUsername(d.CheckUser)
			if err == nil {
				d.CheckUser = "taken"
			}
		} else {

			_, err := data.GetUserByMail(d.CheckMail)
			if err == nil {
				d.CheckMail = "taken"
			}
		}
		a, err := json.Marshal(d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
	}
}
