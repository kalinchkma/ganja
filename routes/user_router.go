package routes

import (
	"ganja/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(actx *interfaces.AppContext, rootRouter *gin.Engine) {
	router := rootRouter.Group("/user")
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "it's working"})
	})

	router.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "wow looks like you are on fire"})
	})
}