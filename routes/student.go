package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/therusetiawan/lks2025cc-backend/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/students", controllers.CreateStudent)
    return r
}
