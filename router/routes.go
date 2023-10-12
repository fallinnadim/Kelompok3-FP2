package router

import (
	controller "fp2/controller/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *controller.AuthenticationController) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})
	authenticationRouter := service.Group("/users")
	authenticationRouter.POST("/register", a.Register)
	authenticationRouter.POST("/login", a.Login)

	return service
}
