package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name string `json:"name"`
}

func CreateProject(c *gin.Context) {

	//bind the request to the struct
	var req CreateProjectRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
	}

	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
		return
	}
	defer db.Close()

	_, err1 := db.Query("INSERT INTO projects (name) VALUES(?)", req.Name)
	if err1 != nil {
		utils.Logger.Err(err1).Msg("Error executing db query")
		c.JSON(http.StatusBadRequest, "Error executing db query")
		return
	}

	c.JSON(http.StatusOK, "Project created successfully!")
}
