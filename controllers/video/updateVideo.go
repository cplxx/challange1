package video

import (
	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func UpdateVideo(ctx *gin.Context) {
	ctx.BindJSON(&models.VideoModel)

	var video models.Video

	result := inits.DB.First(&video, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&video).Updates(models.Video{Title: models.VideoModel.Title, Description: models.VideoModel.Description, Url: models.VideoModel.Url})

	ctx.JSON(200, gin.H{"data": video})
}
