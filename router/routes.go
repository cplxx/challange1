package router

import (
	"github.com/cplxx/AluraChallenge/controllers"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/video", controllers.SearchVideo)
	v1.GET("/videos", controllers.SearchVideoById)
	v1.POST("/video", controllers.CreateVideo)
	v1.PUT("/video", controllers.UpdateVideo)
}
