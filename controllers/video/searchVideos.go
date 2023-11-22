package video

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func SearchVideoById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Esse id não existe")
		return
	}

	video := models.Video{}
	if err := inits.DB.First(&video, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "video não encotrado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": video})
}
