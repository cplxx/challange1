package router

import (
	"github.com/cplxx/AluraChallenge/controllers/category"
	"github.com/cplxx/AluraChallenge/controllers/video"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	vd := router.Group("/vd")
	vd.GET("/video", video.SearchVideo)
	vd.GET("/videos", video.SearchVideoById)
	vd.POST("/video", video.CreateVideo)
	vd.PUT("/video", video.UpdateVideo)
	vd.DELETE("/video", video.DeleteVideo)

	cg := router.Group("cg")
	cg.GET("/category", category.SearchCategory)
	cg.GET("/categorys", category.SearchCategoryById)
	cg.POST("/category", category.CreateCategory)
	cg.PUT("/category", category.UpdateCategory)
	cg.DELETE("/category", category.DeleteCategory)
}
