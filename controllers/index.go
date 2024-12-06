package controllers

import "github.com/gin-gonic/gin"

func IndexRouters(ginEngine *gin.Engine) {
	SetPostRouter(ginEngine)
}
