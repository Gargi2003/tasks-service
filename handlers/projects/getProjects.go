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

// lists down all projects associated with the logged-in user
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
