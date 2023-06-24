package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	Name           string `json:"name"`
	Project_Key    string `json:"project_key"`
	Lead           string `json:"leader"`
	Project_Avatar string `json:"project_avatar"`
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param project body CreateProjectRequest true "Project details"
// @Success 200 {string} string "Project created successfully"
// @Failure 400 {string} string "Error binding req object"
// @Failure 500 {string} string "Error executing db query"
// @Router /projects [post]
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

	_, err1 := db.Query("INSERT INTO projects (name,project_key,leader,project_avatar) VALUES(?, ?, ?, ?)", req.Name, req.Project_Key, req.Lead, req.Project_Avatar)
	if err1 != nil {
		utils.Logger.Err(err1).Msg("Error executing db query")
		c.JSON(http.StatusBadRequest, "Error executing db query")
		return
	}

	c.JSON(http.StatusOK, "Project created successfully!")
}
