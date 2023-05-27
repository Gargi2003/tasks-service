package main

import (
	handler "tasks/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/tasks", handler.CreateTasks)
	router.GET("/tasks/list", handler.ListTasks)
	router.DELETE("/tasks/delete", handler.DeleteTask)
	router.PUT("/tasks/edit", handler.EditTasks)
	router.Run(":8081")
}
