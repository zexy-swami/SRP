package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zexy-swami/SRP/SRP_web/internal/handlers"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("web/static/*.html")
	handlers.InitHandlers(router)
	return router
}
