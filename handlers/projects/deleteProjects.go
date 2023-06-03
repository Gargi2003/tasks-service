package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

// lists down all tasks associated with the loggedin user
func DeleteProject(c *gin.Context) {

	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	projectId := c.Query("id")
	rows, err := db.Exec("delete from projects where id=?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
	}
	result, err := rows.RowsAffected()
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while getting the number of affected rows")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if result == 0 {
		c.JSON(http.StatusNotFound, "no project found with the projectId")
		return
	}

	utils.Logger.Info().Msg("Project deleted successfully!")
	c.JSON(http.StatusOK, "Project Deleted !!!")
}
