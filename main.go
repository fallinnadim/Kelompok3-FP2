package main

import (
	"fp2/config"
	smController "fp2/controller/social_media"
	controller "fp2/controller/users"
	repository "fp2/repository/auth"
	smRepository "fp2/repository/social_media"
	userRepository "fp2/repository/users"
	"fp2/router"
	services "fp2/services/auth"
	smServices "fp2/services/social_media"
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
	socialMediaRepository := smRepository.NewSocialMediaRepositoryImpl(db)

	authenticationService := services.NewAuthServiceImpl(authRepository, validate)
	userService := userServices.NewUserServiceImpl(userRepository, authRepository, validate)
	socialMediaService := smServices.NewSocialMediaServiceImpl(socialMediaRepository, validate)

	authenticationController := controller.NewAuthenticationController(authenticationService)
	userController := controller.NewUserController(userService)
	socialMediaController := smController.NewSocialMediaController(socialMediaService)

	routes := router.NewRouter(userRepository, socialMediaRepository, authenticationController, userController, socialMediaController)

	server := &http.Server{
		Addr:    ":3030",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
