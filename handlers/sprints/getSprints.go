package handlers

import (
	"net/http"
	utils "tasks/common"

	"github.com/gin-gonic/gin"
)

type GetSprint struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	ProjectID int    `json:"project_id"`
}

// lists down all sprints associated with the logged-in user

// ListSprints godoc
// @Summary List all sprints
// @Description Lists down all sprints associated with the logged-in user
// @Tags Sprints
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} GetSprint
// @Failure 500 {string} string "Internal Server Error"
// @Router /sprints/list [get]
func ListSprints(c *gin.Context) {

	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var sprints []GetSprint
	rows, err := db.Query("SELECT id, name, start_date, end_date, project_id from sprints")
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		sprint := GetSprint{}
		var start, end []uint8

		err := rows.Scan(&sprint.ID, &sprint.Name, &start, &end, &sprint.ProjectID)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}
		sprint.StartDate = string(start)
		sprint.EndDate = string(end)

		sprints = append(sprints, sprint)
	}

	c.JSON(http.StatusOK, sprints)
}

// GetSprintById godoc
// @Summary Get a sprint by ID
// @Description Get a sprint by its ID
// @Tags Sprints
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Sprint ID"
// @Success 200 {array} GetSprint
// @Failure 500 {string} string "Internal Server Error"
// @Router /sprints/get [get]
func GetSprintById(c *gin.Context) {

	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var sprints []GetSprint
	sprintId := c.Query("id")
	rows, err := db.Query("SELECT id, name, start_date, end_date, project_id from sprints where id=?", sprintId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		sprint := GetSprint{}
		err := rows.Scan(&sprint.ID, &sprint.Name, &sprint.StartDate, &sprint.EndDate, &sprint.ProjectID)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}

		sprints = append(sprints, sprint)
	}

	c.JSON(http.StatusOK, sprints)
}

// GetSprintByProjectId godoc
// @Summary Get a sprint by ID
// @Description Get a sprint by its ID
// @Tags Sprints
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Sprint ID"
// @Success 200 {array} GetSprint
// @Failure 500 {string} string "Internal Server Error"
// @Router /sprints/get [get]
func GetSprintByProjectId(c *gin.Context) {

	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var sprints []GetSprint
	projectId := c.Query("id")
	rows, err := db.Query("SELECT id, name, start_date, end_date, project_id from sprints where project_id=?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		sprint := GetSprint{}
		err := rows.Scan(&sprint.ID, &sprint.Name, &sprint.StartDate, &sprint.EndDate, &sprint.ProjectID)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}

		sprints = append(sprints, sprint)
	}

	c.JSON(http.StatusOK, sprints)
}

// GetSprintsByProjectName godoc
// @Summary Get a sprint by Project Name
// @Description Get a sprint by Project Name
// @Tags Sprints
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id query string true "Project Name"
// @Success 200 {array} GetSprint
// @Failure 500 {string} string "Internal Server Error"
// @Router /sprints/get [get]
func GetSprintsByProjectName(c *gin.Context) {
	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	var sprints []GetSprint
	proj_name := c.Query("project_name")
	rows, err := db.Query("select * from sprints where project_id= (select id from projects where name = ?)", proj_name)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occured while executing query")
		c.JSON(http.StatusInternalServerError, "Error occured while executing query")
	}
	for rows.Next() {
		sprint := GetSprint{}
		err := rows.Scan(&sprint.ID, &sprint.Name, &sprint.StartDate, &sprint.EndDate, &sprint.ProjectID)
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while populating struct with db data")
			c.JSON(http.StatusInternalServerError, "Error occurred while populating struct with db data")
		}
		sprints = append(sprints, sprint)
	}
	c.JSON(http.StatusOK, sprints)
}
