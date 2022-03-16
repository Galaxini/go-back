package controllers

import (
	"encoding/json"
	"fmt"
	"go-back/app/models"
	"go-back/app/services"
	"net/http"
)

func AddItemHandler(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("token") == "" {
		fmt.Fprintf(w, "Please enter a valid token.\r\n")
		return
	}
	if req.FormValue("title") == "" {
		fmt.Fprintf(w, "Please enter a valid title.\r\n")
		return
	}
	if req.FormValue("description") == "" {
		fmt.Fprintf(w, "Please enter a valid description.\r\n")
		return
	}
	if req.FormValue("price") == "" {
		fmt.Fprintf(w, "Please enter a valid price.\r\n")
		return
	}

	user_id, err := services.GetUserID(req.FormValue("token"))

	response := "User not found"

	if user_id != 0 {
		response, err = services.AddItem(user_id, req.FormValue("title"), req.FormValue("description"), req.FormValue("price"))
	}

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, response)
	}
}

func GetItemsHandler(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("token") == "" {
		fmt.Fprintf(w, "Please enter a valid token.\r\n")
		return
	}

	user_id, err := services.GetUserID(req.FormValue("token"))

	response := models.Item{}

	if user_id != 0 {
		response, err = services.GetItems(user_id)
	}
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

func EditItemsHandler(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("token") == "" {
		fmt.Fprintf(w, "Please enter a valid token.\r\n")
		return
	}
	if req.FormValue("title") == "" {
		fmt.Fprintf(w, "Please enter a valid title.\r\n")
		return
	}
	if req.FormValue("description") == "" {
		fmt.Fprintf(w, "Please enter a valid description.\r\n")
		return
	}
	if req.FormValue("price") == "" {
		fmt.Fprintf(w, "Please enter a valid price.\r\n")
		return
	}
	if req.FormValue("id") == "" {
		fmt.Fprintf(w, "Please enter a valid id.\r\n")
		return
	}
	user_id, err := services.GetUserID(req.FormValue("token"))
	response := "User not found"
	if user_id != 0 {
		response, err = services.EditItems(req.FormValue("title"), req.FormValue("description"), req.FormValue("price"), req.FormValue("id"))
	}
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, response)
	}
}

func DeleteItemsHandler(w http.ResponseWriter, req *http.Request) {
	if req.FormValue("token") == "" {
		fmt.Fprintf(w, "Please enter a valid token.\r\n")
		return
	}
	if req.FormValue("id") == "" {
		fmt.Fprintf(w, "Please enter a valid id.\r\n")
		return
	}
	user_id, err := services.GetUserID(req.FormValue("token"))
	response := "User not found"
	if user_id != 0 {
		response, err = services.DeleteItems(req.FormValue("id"))
	}
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, response)
	}
}
