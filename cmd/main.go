package main

import (
	"fmt"
	"net/http"
	"os"

	cors "github.com/rs/cors"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}

	//Cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "*"},
		AllowCredentials: true,
	})

	if err := http.ListenAndServe(":5000", c.Handler(handler.Router)); err != nil {
		fmt.Println("Failed to set up server")
	}
}
