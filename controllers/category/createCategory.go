package category

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context) {

	if err := ctx.BindJSON(models.Category{}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	result := inits.DB.Create(models.Category{})
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category created successfully"})
}
