package main

import (
	"net/http"

	"go-back/app/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views/")))

	router.HandleFunc("/auth/register", controllers.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth/login", controllers.LoginHandler).Methods("POST")
	router.HandleFunc("/item/addItems", controllers.AddItemHandler).Methods("POST")
	router.HandleFunc("/item/getItems", controllers.GetItemsHandler).Methods("GET")
	router.HandleFunc("/item/editItems", controllers.EditItemsHandler).Methods("POST")
	router.HandleFunc("/item/deleteItems", controllers.DeleteItemsHandler).Methods("POST")

	http.ListenAndServe(":8090", router)
}
