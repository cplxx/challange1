package category

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func SearchCategoryById(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "esse id não existe")
		return
	}

	category := models.Category{}
	if err := inits.DB.First(&category, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "categoria não encotrada"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"category": category})
}
