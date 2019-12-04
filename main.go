package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-http-server/config"
	"golang-http-server/controller"
	"golang-http-server/models"
	"net/http"
	"os"
)

func main() {
	db := config.Init()
	db.Debug().AutoMigrate(&models.Employee{})
	router := mux.NewRouter()
	router.HandleFunc("/api/employee/create", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee/getAll", controller.GetAllEmployees).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)

	error := http.ListenAndServe(":"+port, router)
	if error != nil {
		fmt.Print(error)
	}

}
