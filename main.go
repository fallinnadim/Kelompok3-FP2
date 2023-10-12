package main

import (
	"fp2/config"
	controller "fp2/controller/users"
	repository "fp2/repository/auth"
	userRepository "fp2/repository/users"
	"fp2/router"
	services "fp2/services/auth"
	userServices "fp2/services/users"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

func main() {
	// Database
	db := config.StartDB()
	validate := validator.New()

	authRepository := repository.NewAuthRepositoryImpl(db)
	userRepository := userRepository.NewUserRepositoryImpl(db)
	authenticationService := services.NewAuthServiceImpl(authRepository, validate)
	userService := userServices.NewUserServiceImpl(userRepository, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService)
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userRepository, authenticationController, userController)

	server := &http.Server{
		Addr:    ":3030",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
