package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mkilic91/goBlog/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.GetIndex)
	router.GET("/p/:page", controllers.GetIndex)
	router.GET("/c/:slug", controllers.GetPost)
}
