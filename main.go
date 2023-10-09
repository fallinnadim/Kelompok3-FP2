package main

import (
	"fp2/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	config.StartDB()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	server := &http.Server{
		Addr:    ":3030",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
