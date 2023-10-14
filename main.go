package main

import (
	"fp2/config"
	cController "fp2/controller/comment"
	pController "fp2/controller/photo"
	smController "fp2/controller/social_media"
	controller "fp2/controller/users"
	repository "fp2/repository/auth"
	commentRepository "fp2/repository/comment"
	photoRepository "fp2/repository/photo"
	smRepository "fp2/repository/social_media"
	userRepository "fp2/repository/users"
	"fp2/router"
	services "fp2/services/auth"
	cServices "fp2/services/comment"
	pServices "fp2/services/photo"
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
	photoRepository := photoRepository.NewPhotoRepositoryImpl(db)
	socialMediaRepository := smRepository.NewSocialMediaRepositoryImpl(db)
	commentRepostiory := commentRepository.NewCommentRepositoryImpl(db)

	authenticationService := services.NewAuthServiceImpl(authRepository, validate)
	userService := userServices.NewUserServiceImpl(userRepository, authRepository, validate)
	socialMediaService := smServices.NewSocialMediaServiceImpl(socialMediaRepository, validate)
	photoService := pServices.NewPhotoServiceImpl(photoRepository, validate)
	commentService := cServices.NewCommentServiceImpl(commentRepostiory, photoRepository, validate)

	authenticationController := controller.NewAuthenticationController(authenticationService)
	userController := controller.NewUserController(userService)
	socialMediaController := smController.NewSocialMediaController(socialMediaService)
	photoController := pController.NewPhotoController(photoService)
	commentController := cController.NewCommentController(commentService)

	routes := router.NewRouter(userRepository, socialMediaRepository, photoRepository, commentRepostiory, authenticationController, userController, socialMediaController, photoController, commentController)

	server := &http.Server{
		Addr:    ":3030",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
