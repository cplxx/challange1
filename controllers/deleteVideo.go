package controllers

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func DeleteVideo(ctx *gin.Context) {

	inits.DB.Delete(&models.Video{}, ctx.Query("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"data": "video deletado com sucesso"})
}
