package category

import (
	"net/http"

	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/gin-gonic/gin"
)

func DeleteCategory(ctx *gin.Context) {
	inits.DB.Delete(&models.Category{}, ctx.Query("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"data": "categoria deletada com sucesso"})
}
