package controllers

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func SearchCategory(ctx *gin.Context) {
	var category []models.Category

	result := inits.DB.Find(category)
	if result == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
