package controllers

import "github.com/gin-gonic/gin"

func IndexRouters() *gin.Engine {
	router := gin.Default()

	SetPostRouter(router)

	return router
}
