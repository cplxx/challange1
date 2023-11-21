package controllers

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func DeleteVideo(ctx *gin.Context) {
	id := ctx.Param("id")

	inits.DB.Delete(&models.Video{}, id)
	ctx.JSON(http.StatusOK, gin.H{"sucess": "video deletado com sucesso"})
}
