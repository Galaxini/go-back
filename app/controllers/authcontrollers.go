package controllers

import (
	"fmt"
	"net/http"
	"timesavvy/app/services"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {

	// validation
	if req.FormValue("name") == "" {
		fmt.Fprintf(w, "Please enter a valid name.\r\n")
		return
	}
	if req.FormValue("password") == "" {
		fmt.Fprintf(w, "Please enter a valid password.\r\n")
		return
	}
	if req.FormValue("email") == "" {
		fmt.Fprintf(w, "Please enter a valid email.\r\n")
		return
	}

	response, err := services.RegisterUser(req.FormValue("name"), req.FormValue("password"), req.FormValue("email"))

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, response)
	}
}
