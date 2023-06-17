package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type CreateSprintRequest struct {
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	ProjectID int    `json:"project_id"`
}

// CreateSprint godoc
// @Summary Create a new sprint
// @Description Create a new sprint for a project
// @Tags Sprints
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param req body CreateSprintRequest true "Sprint details"
// @Success 200 {string} string "Sprint created successfully!"
// @Failure 400 {string} string "Error executing db query"
// @Failure 500 {string} string "Internal Server Error"
// @Router /sprints [post]
func CreateSprint(c *gin.Context) {

	//bind the request to the struct
	var req CreateSprintRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusInternalServerError, "Error binding req object")
	}

	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	_, err1 := db.Query("INSERT INTO sprints (name, start_date, end_date, project_id) VALUES (?, ?, ?, ?)", req.Name, req.StartDate, req.EndDate, req.ProjectID)
	if err1 != nil {
		utils.Logger.Err(err).Msg("Error executing db query")
		c.JSON(http.StatusBadRequest, "Error executing db query")
	}

	c.JSON(http.StatusOK, "Sprint created successfully!")
}
