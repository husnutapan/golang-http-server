package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	pass = "root"
	name = "golang"
)

func Init() {
	var db *sql.DB
	var err error

	config := initialDbConfig()
	sqlFormat := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[host], config[port],
		config[user], config[pass], config[name])

	db, err = sql.Open("postgres", sqlFormat)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully ping to database")
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
