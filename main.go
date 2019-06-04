package main

import (
	"golang_demo/db"
	"golang_demo/router"
)

func main()  {
	dbConfig := db.MyDbConfig()
	db.CreateDB(&dbConfig)

	router.StartRouter()
}
