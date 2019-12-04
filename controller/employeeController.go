package controller

import (
	"encoding/json"
	"fmt"
	"golang-http-server/models"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := &models.Employee{}
	json.NewDecoder(r.Body).Decode(employee)
	createdEmployee := employee.CreateEmployee()
	json.NewEncoder(w).Encode(createdEmployee)

}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
