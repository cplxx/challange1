package video

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func CreateVideo(ctx *gin.Context) {
	var input models.Video

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video := models.Video{
		Title:       input.Title,
		Description: input.Description,
		Url:         input.Url,
	}

	result := inits.DB.Create(&video)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video created successfully"})
}
