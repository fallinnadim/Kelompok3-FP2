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
				"message": "Silahkan login terlebih dahulu",
			})
			return
		}
		config.LoadConfig()
		sub, err := utils.ValidateToken(token, os.Getenv("TOKEN_SECRET"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Token tidak valid",
			})
			return
		}
		id, errId := strconv.Atoi(fmt.Sprint(sub))
		if errId != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Token tidak valid",
			})
			return
		}
		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "User dengan token ini sudah tidak ada lagi",
			})
			return
		}
		ctx.Set("userId", result.Id)
		ctx.Next()
	}
}
