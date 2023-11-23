package category

import (
	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func UpdateCategory(ctx *gin.Context) {
	ctx.BindJSON(&models.CategoryModel)

	var category models.Category

	result := inits.DB.First(&category, ctx.Query("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&category).Updates(models.Category{Title: models.CategoryModel.Title, Color: models.CategoryModel.Color})

	ctx.JSON(200, gin.H{"category": category})
}
