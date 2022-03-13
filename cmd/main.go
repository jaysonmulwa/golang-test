package main

import (
	"fmt"
	"net/http"

	customers "github.com/jaysonmulwa/jumia/internal/customer"
	database "github.com/jaysonmulwa/jumia/internal/database"
	handler "github.com/jaysonmulwa/jumia/internal/handler"
)

func main() {
	//Init New database connection
	db, _ := database.Connect()

	//Init New customer service
	newCustomerService := customers.NewCustomerService(db)

	//Init New routes handler
	handler := handler.NewHandler(newCustomerService)

	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
	}
}

/*
//Add Gorm ORM
//Add all Routes
//Add logic for routes
//Add tests

//!Dockerize, deploy to Heroku/DigitalOceam
//!Vue frontend and host on netlify
//!Submit with deployment

//!Refactor code for testability
//!Add more testa



valid - ok, Nok, all*/
