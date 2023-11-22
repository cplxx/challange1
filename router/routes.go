package router

import (
	"github.com/cplxx/AluraChallenge/controllers"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	vd := router.Group("/vd")
	vd.GET("/video", controllers.SearchVideo)
	vd.GET("/videos", controllers.SearchVideoById)
	vd.POST("/video", controllers.CreateVideo)
	vd.PUT("/video", controllers.UpdateVideo)
	vd.DELETE("/video", controllers.DeleteVideo)

	cg := router.Group("cg")
	cg.GET("/category", controllers.SearchCategory)
	cg.POST("/category", controllers.CreateCategory)
}
