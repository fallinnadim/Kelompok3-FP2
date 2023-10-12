package middleware

import (
	"fmt"
	"fp2/config"
	repository "fp2/repository/users"

	"fp2/utils"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializedUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "please login first",
			})
			return
		}
		config.LoadConfig()
		sub, err := utils.ValidateToken(token, os.Getenv("TOKEN_SECRET"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Invalid token",
			})
			return
		}
		id, errId := strconv.Atoi(fmt.Sprint(sub))
		if errId != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Invalid token",
			})
			return
		}
		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "the user belonging to this token no longer exist",
			})
			return
		}
		ctx.Set("userId", result.Id)
		ctx.Next()
	}
}
