package main

import (
	_ "tasks/docs"
	auditHandler "tasks/handlers/audits"
	projectHandler "tasks/handlers/projects"
	sprintHandler "tasks/handlers/sprints"
	taskhandler "tasks/handlers/tasks"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Tasks Service
// @description Tasks API in go using gin-framework
// @version 1.0
// @host localhost:8081
// @BasePath /api
func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//tasks api
	router.POST("/tasks", taskhandler.CreateTasks)
	router.GET("/tasks/list", taskhandler.ListTasks)
	router.GET("/tasks/get", taskhandler.GetTaskByUseridAndTaskId)
	router.DELETE("/tasks/delete", taskhandler.DeleteTask)
	router.PUT("/tasks/edit", taskhandler.EditTasks)
	router.GET("/tasks/getByProjectId", taskhandler.GetTaskByProjectId)
	router.GET("/tasks/getBySprintId", taskhandler.GetTaskBySprintIdandProjectId)
	router.GET("/tasks/getByUserId", taskhandler.GetTaskByUserId)

	//sprints api
	router.POST("/sprints", sprintHandler.CreateSprint)
	router.GET("/sprints/list", sprintHandler.ListSprints)
	router.GET("/sprints/get", sprintHandler.GetSprintById)
	router.DELETE("/sprints/delete", sprintHandler.DeleteSprint)
	router.PUT("/sprints/edit", sprintHandler.EditSprint)
	router.GET("/sprints/getByProjectId", sprintHandler.GetSprintByProjectId)
	router.GET("/sprints/getByProjectName", sprintHandler.GetSprintsByProjectName)

	//projects api
	router.POST("/projects", projectHandler.CreateProject)
	router.GET("/projects/list", projectHandler.ListProjects)
	router.GET("/projects/get", projectHandler.GetProjectById)
	router.DELETE("/projects/delete", projectHandler.DeleteProject)
	router.PUT("/projects/edit", projectHandler.EditProject)

	//audits api
	router.GET("/audit/list", auditHandler.ListAudits)

	router.Run(":8081")
}
