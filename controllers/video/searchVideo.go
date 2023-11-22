package video

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func SearchVideo(ctx *gin.Context) {
	var videos []models.Video

	result := inits.DB.Find(&videos)
	if result == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": videos})
}
