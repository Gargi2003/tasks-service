package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type GetProject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// lists down all projects associated with the logged-in user
// ListProjects godoc
// @Summary List all projects
// @Description Get a list of all projects associated with the logged-in user
// @Tags Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} []GetProject
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects [get]
func ListProjects(c *gin.Context) {

	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var projects []GetProject
	rows, err := db.Query("SELECT id, name from projects")
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		project := GetProject{}

		err := rows.Scan(&project.ID, &project.Name)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}

		projects = append(projects, project)
	}

	c.JSON(http.StatusOK, projects)
}

// GetProjectById godoc
// @Summary Get a project by ID
// @Description Get a project by its ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Project ID"
// @Success 200 {object} GetProject
// @Failure 404 {string} string "No project found with the project ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/{id} [get]
func GetProjectById(c *gin.Context) {

	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	projectId := c.Query("id")
	rows, err := db.Query("SELECT name from projects where id=?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	project := GetProject{}
	err1 := rows.Scan(&project.Name)
	if err1 != nil {
		utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
	}

	c.JSON(http.StatusOK, project)
}
