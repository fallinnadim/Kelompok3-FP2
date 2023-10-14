package router

import (
	cController "fp2/controller/comment"
	pController "fp2/controller/photo"
	smController "fp2/controller/social_media"
	controller "fp2/controller/users"
	"fp2/middleware"
	crRepository "fp2/repository/comment"
	pRepository "fp2/repository/photo"
	srRepository "fp2/repository/social_media"
	repository "fp2/repository/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(ur repository.UserRepository, sr srRepository.SocialMediaRepository, pr pRepository.PhotoRepository, cr crRepository.CommentRepository, a *controller.AuthenticationController, u *controller.UserController, sm *smController.SocialMediaController, p *pController.PhotoController, c *cController.CommentController) *gin.Engine {
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
	commentRouter := service.Group("/comments")
	commentRouter.Use(middleware.AuthenticatedUser(ur))
	{
		commentRouter.POST("", c.CreateComment)
		commentRouter.GET("", c.GetAllComment)
		commentRouter.PUT("/:commentId", middleware.AuthorizedUserC(cr), c.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.AuthorizedUserC(cr), c.DeleteComment)
	}

	return service
}
