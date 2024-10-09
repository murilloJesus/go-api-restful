package routes

import (
	"restful/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	return router
}
