package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"pixkeys/infra"
	"pixkeys/infra/middlewares"
	"pixkeys/routes"
	"time"
)

func main() {
	//Initializing dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//SETUP DB
	if err := infra.SetupDB(); err != nil {
		log.Fatal("[DB Setup Error]", err)
	}

	//SETUP ROUTES
	router := mux.NewRouter()
	router.Use(middlewares.Logging)

	router.HandleFunc("/pix_keys", routes.CreatePixKey).Methods("POST")

	//START SERVER
	address := ":3001"
	server := http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Running server on port ", address)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error running server: ", err)
	}
}
