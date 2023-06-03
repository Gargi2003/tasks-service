package main

import (
	projectHandler "tasks/handlers/projects"
	sprintHandler "tasks/handlers/sprints"
	taskhandler "tasks/handlers/tasks"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//tasks api
	router.POST("/tasks", taskhandler.CreateTasks)
	router.GET("/tasks/list", taskhandler.ListTasks)
	router.DELETE("/tasks/delete", taskhandler.DeleteTask)
	router.PUT("/tasks/edit", taskhandler.EditTasks)

	//sprints api
	router.POST("/sprints", sprintHandler.CreateSprint)
	router.GET("/sprints/list", sprintHandler.ListSprints)
	router.DELETE("/sprints/delete", sprintHandler.DeleteSprint)
	router.PUT("/sprints/edit", sprintHandler.EditSprint)

	//projects api
	router.POST("/projects", projectHandler.CreateProject)
	router.GET("/projects/list", projectHandler.ListProjects)
	router.DELETE("/projects/delete", projectHandler.DeleteProject)
	router.PUT("/projects/edit", projectHandler.EditProject)

	router.Run(":8081")
}
