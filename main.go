package main

import (
	"go-crud/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main()  {

	log.Info().Msg("Starter Server")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, "welcone home")
	})
	server := &http.Server{
		Addr: ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}