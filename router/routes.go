package router

import (
	smController "fp2/controller/social_media"
	controller "fp2/controller/users"
	"fp2/middleware"
	repository "fp2/repository/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(ur repository.UserRepository, a *controller.AuthenticationController, u *controller.UserController, sm *smController.SocialMediaController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})
	authenticationRouter := service.Group("/users")
	authenticationRouter.POST("/register", a.Register)
	authenticationRouter.POST("/login", a.Login)
	authenticationRouter.Use(middleware.DeserializedUser(ur))
	{
		authenticationRouter.PUT("", u.UpdateUser)
		authenticationRouter.DELETE("", u.DeleteUser)
	}
	socialMediaRouter := service.Group("/socialmedias")
	socialMediaRouter.Use(middleware.DeserializedUser(ur))
	{
		socialMediaRouter.POST("", sm.CreateSocialMedia)
	}

	return service
}
