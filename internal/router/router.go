package router

import (
	"github.com/geekswamp/zen/internal/di"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	apiV1 := engine.Group("/api/v1")

	userGroup := apiV1.Group("/user")

	userHandler := di.InitUserHandler()
	userGroup.POST("/register", userHandler.Register)
	userGroup.GET("/detail/:id", userHandler.GetDetail)
	userGroup.DELETE("/delete/:id", userHandler.HardDelete)
	userGroup.PATCH("/mark-delete/:id", userHandler.SoftDelete)
	userGroup.PATCH("/set-active/:id", userHandler.SetToActive)
	userGroup.PATCH("/set-inactive/:id", userHandler.SetToInactive)
}
