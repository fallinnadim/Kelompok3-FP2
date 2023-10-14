package router

import (
	pController "fp2/controller/photo"
	smController "fp2/controller/social_media"
	controller "fp2/controller/users"
	"fp2/middleware"
	pRepository "fp2/repository/photo"
	srRepository "fp2/repository/social_media"
	repository "fp2/repository/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(ur repository.UserRepository, sr srRepository.SocialMediaRepository, pr pRepository.PhotoRepository, a *controller.AuthenticationController, u *controller.UserController, sm *smController.SocialMediaController, p *pController.PhotoController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})
	authenticationRouter := service.Group("/users")
	authenticationRouter.POST("/register", a.Register)
	authenticationRouter.POST("/login", a.Login)
	authenticationRouter.Use(middleware.AuthenticatedUser(ur))
	{
		authenticationRouter.PUT("", u.UpdateUser)
		authenticationRouter.DELETE("", u.DeleteUser)
	}
	socialMediaRouter := service.Group("/socialmedias")
	socialMediaRouter.Use(middleware.AuthenticatedUser(ur))
	{
		socialMediaRouter.POST("", sm.CreateSocialMedia)
		socialMediaRouter.GET("", sm.GetAllSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middleware.AuthorizedUserSm(sr), sm.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.AuthorizedUserSm(sr), sm.DeleteSocialMedia)
	}
	photoRouter := service.Group("/photos")
	photoRouter.Use(middleware.AuthenticatedUser(ur))
	{
		photoRouter.POST("", p.CreatePhoto)
		photoRouter.GET("", p.GetAllPhoto)
		photoRouter.PUT("/:photoId", middleware.AuthorizedUserP(pr), p.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.AuthorizedUserP(pr), p.DeletePhoto)
	}

	return service
}
