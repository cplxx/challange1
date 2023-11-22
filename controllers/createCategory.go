package controllers

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result := inits.DB.Create(&category)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category created successfully"})
}
