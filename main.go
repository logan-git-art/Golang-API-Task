package main

import (
	"moule2project/controllers"
	"moule2project/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTasks)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PUT("/tasks/:id", controllers.UpdateTasks)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.Run()
}
