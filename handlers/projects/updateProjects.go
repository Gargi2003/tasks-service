package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type UpdateProjectRequest struct {
	Name string `json:"name"`
}

// EditProject godoc
// @Summary Update a project
// @Description Update a project by its ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Project ID"
// @Param body body UpdateProjectRequest true "Update project request body"
// @Success 200 {string} string "Project updated successfully"
// @Failure 400 {string} string "Error binding req object"
// @Failure 500 {string} string "Error updating project"
// @Router /projects/{id} [put]
func EditProject(c *gin.Context) {
	// Connect to the database
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	// Bind the request to the struct
	var req UpdateProjectRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
		return
	}

	// Build the update query dynamically based on non-empty fields
	updateQuery := "UPDATE projects SET"
	var queryParams []interface{}

	if req.Name != "" {
		updateQuery += " name=?,"
		queryParams = append(queryParams, req.Name)
	}

	updateQuery = updateQuery[:len(updateQuery)-1]

	// Add the WHERE condition
	updateQuery += " WHERE id=?"

	projectId := c.Query("id")
	// Add the user ID parameter to the query parameters
	queryParams = append(queryParams, projectId)

	// Execute the update query
	_, err1 := db.Query(updateQuery, queryParams...)
	if err1 != nil {
		utils.Logger.Err(err1).Msg("Error executing update query")
		c.JSON(http.StatusInternalServerError, "Error updating project")
		return
	}

	// Handle the response
	c.JSON(http.StatusOK, "Project updated successfully")
}
