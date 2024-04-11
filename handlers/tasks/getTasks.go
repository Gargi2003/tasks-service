package handlers

import (
	"fmt"
	"net/http"
	utils "tasks/common"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ListTasks godoc
// @Summary List tasks
// @Description List all tasks for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} utils.Task "List of tasks"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/list [get]
func ListTasks(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	fmt.Println("tokenstring", tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})
	if token == nil {
		utils.Logger.Info().Msg("token is nil")
		return
	}
	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task
	// Extract the user ID associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]

		rows, err := db.Query("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points,reporter, comments, status FROM tasks WHERE user_id=?", userId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while executing query")
			c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
			return
		}

		for rows.Next() {
			task := utils.Task{}
			var createdAt, updatedAt string
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectID, &task.Points, &task.Reporter, &task.Comments, &task.Status)
			if err != nil {
				utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
			}

			task.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
			task.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

			tasks = append(tasks, task)
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByUseridAndTaskId godoc
// @Summary Get tasks
// @Description Get all tasks for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} utils.Task "Fetch tasks"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/get [get]
func GetTaskByUseridAndTaskId(c *gin.Context) {
	// Connect to the db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	taskId := c.Query("taskid")
	userId := c.Query("userid")
	row := db.QueryRow("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points, reporter, comments, status FROM tasks WHERE user_id=? AND id=?", userId, taskId)

	task := utils.Task{}
	var createdAt, updatedAt string
	err1 := row.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectID, &task.Points, &task.Reporter, &task.Comments, &task.Status)
	if err1 != nil {

		utils.Logger.Err(err1).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")

		return
	}

	task.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	task.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	c.JSON(http.StatusOK, task)

}

// GetTaskByProjectId godoc
// @Summary List tasks
// @Description List all tasks for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} utils.Task "List of tasks"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/list [get]
func GetTaskByProjectId(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task
	projectId := c.Query("projectId")

	rows, err := db.Query("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points,reporter, comments, status FROM tasks WHERE  project_id=?", projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		task := utils.Task{}
		var createdAt, updatedAt string
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectID, &task.Points, &task.Reporter, &task.Comments, &task.Status)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}

		task.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		task.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskBySprintId godoc
// @Summary List tasks
// @Description List all tasks for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} utils.Task "List of tasks"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/list [get]
func GetTaskBySprintIdandProjectId(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task
	sprintId := c.Query("sprintId")
	projectId := c.Query("projectId")
	rows, err := db.Query("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points,reporter, comments, status FROM tasks WHERE  sprint_id=? and project_id=?", sprintId, projectId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}

	for rows.Next() {
		task := utils.Task{}
		var createdAt, updatedAt string
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectID, &task.Points, &task.Reporter, &task.Comments, &task.Status)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}

		task.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		task.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByUserId godoc
// @Summary List tasks
// @Description List all tasks for the logged-in user
// @Tags Tasks
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} utils.Task "List of tasks"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/list [get]
func GetTaskByUserId(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})

	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var tasks []utils.Task

	// Extract the user ID associated with the current session
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["sub"]
		rows, err := db.Query("SELECT id, title, description, created_at, updated_at, user_id, issue_type, assignee, sprint_id, project_id, points,reporter, comments, status FROM tasks WHERE  user_id=?", userId)
		if err != nil {
			utils.Logger.Err(err).Msg("Error occurred while executing query")
			c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
			return
		}

		for rows.Next() {
			task := utils.Task{}
			var createdAt, updatedAt string
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &createdAt, &updatedAt, &task.UserID, &task.IssueType, &task.Assignee, &task.Sprint, &task.ProjectID, &task.Points, &task.Reporter, &task.Comments, &task.Status)
			if err != nil {
				utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
			}

			task.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
			task.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

			tasks = append(tasks, task)
		}
	}
	c.JSON(http.StatusOK, tasks)
}
