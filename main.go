package main

import (
	handler "tasks/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/tasks", handler.CreateTasks)
	router.GET("/tasks/list", handler.ListTasks)
	router.Run(":8081")
}
