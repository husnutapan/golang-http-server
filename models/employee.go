package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang-http-server/config"
)

type Employee struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (e *Employee) ValidationModel() map[string]string {
	validations := make(map[string]string)

	if e.Name == "" {
		validations["Name"] = "Name cannot empty value"
	}

	if e.Surname == "" {
		validations["Surname"] = "Surname cannot empty value"
	}

	if e.Age <= 0 {
		validations["Age"] = "Age cannot lower than zero"
	}
	return validations
}

func (e *Employee) CreateEmployee() bool {
	validateMessages := e.ValidationModel()

	if len(validateMessages) > 0 {
		for key, value := range validateMessages {
			fmt.Println("Key:"+key, " Value:"+value)
		}
	}
	DB().Create(e)
	return true
}

func DB() *gorm.DB {
	return config.Init()
}
