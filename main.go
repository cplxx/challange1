package main

import (
	"github.com/cplxx/AluraChallenge/inits"
	"github.com/cplxx/AluraChallenge/models"
	"github.com/cplxx/AluraChallenge/router"
	"github.com/gin-gonic/gin"
)

func main() {
	inits.LoadEnv()
	inits.DBInit()
	inits.DB.AutoMigrate(&models.Video{})
	r := gin.Default()

	router.Initialize()

	r.Run()
}
