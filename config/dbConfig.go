package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	pass = "root"
	name = "golang"
)

var db *gorm.DB

func Init() *gorm.DB {
	var err error

	config := initialDbConfig()
	sqlFormat := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[host], config[port],
		config[user], config[pass], config[name])

	conn, err := gorm.Open("postgres", sqlFormat)
	if err != nil {
		panic(err)
	}
	db = conn
	fmt.Println("Successfully ping to database")
	return db
}

func initialDbConfig() map[string]string {
	configs := make(map[string]string)
	configs[host] = host
	configs[port] = port
	configs[user] = user
	configs[pass] = pass
	configs[name] = name
	return configs
}
