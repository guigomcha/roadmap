package main

import (
	"db-manager/pkg/connection"
	"db-manager/pkg/models"
	"log"
)

func main() {
	db, err := connection.ConnectToPostgreSQL()
	if err != nil {
		log.Fatal(err)
	}

	// Perform database migration
	err = db.AutoMigrate(&models.Spot{})
	if err != nil {
		log.Fatal(err)
	}

	// Your CRUD operations go here
}
