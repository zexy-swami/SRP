package handlers

import "github.com/gin-gonic/gin"

func InitHandlers(router *gin.Engine) {
	router.NoRoute(methodNotAllowedHandler)
	router.GET("/", rootHandler)
	router.GET("/sign-up", signUpHandler)
	router.POST("/verification", verificationHandler)
}
