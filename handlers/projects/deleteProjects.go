package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Project ID"
// @Success 200 {string} string "Project Deleted !!!"
// @Failure 404 {string} string "No project found with the project ID"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/delete [delete]
func DeleteProject(c *gin.Context) {
	// Connect to the database
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer db.Close()

	projectId := c.Query("id")

	// Disable foreign key checks temporarily
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS=0")
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while disabling foreign key checks")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Delete the tasks associated with the project
	_, err = db.Exec("DELETE FROM tasks WHERE project_id = ?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while deleting tasks")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Enable foreign key checks again
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS=1")
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while enabling foreign key checks")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Delete the sprints associated with the project
	_, err = db.Exec("DELETE FROM sprints WHERE project_id = ?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while deleting sprints")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Delete the project
	result, err := db.Exec("DELETE FROM projects WHERE id = ?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while deleting project")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while getting the number of affected rows")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, "No project found with the provided projectId")
		return
	}

	utils.Logger.Info().Msg("Project, associated sprints, and tasks deleted successfully!")
	c.JSON(http.StatusOK, "Project Deleted !!!")
}
