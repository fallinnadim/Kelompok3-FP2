package main

import (
	"fp2/config"
	controller "fp2/controller/users"
	repository "fp2/repository/auth"
	"fp2/router"
	services "fp2/services/auth"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func main() {
	// Database
	db := config.StartDB()
	validate := validator.New()

	userRepository := repository.NewAuthRepositoryImpl(db)
	authenticationService := services.NewAuthServiceImpl(userRepository, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService)

	routes := router.NewRouter(authenticationController)

	server := &http.Server{
		Addr:    ":3030",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
