package controller

import (
	"fmt"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
